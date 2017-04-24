package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

type joke struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}

func dataSource() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"joke": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {
	id, ok := d.GetOk("id")
	url := "http://api.icndb.com/jokes/random"
	if ok {
		url = fmt.Sprintf("http://api.icndb.com/jokes/%d", id)
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	j := new(joke)
	json.NewDecoder(resp.Body).Decode(j)

	d.Set("joke", j.Value.Joke)

	d.SetId("-")

	return nil
}
