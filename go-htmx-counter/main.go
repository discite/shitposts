package main

import (
	"fmt"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
	"github.com/chasefleming/elem-go/styles"
	"github.com/savsgio/atreugo/v11"
)

func main() {
	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)

	var count int

	server.POST("/increment", func(rc *atreugo.RequestCtx) error {
		count++
		return rc.TextResponse(fmt.Sprintf("%d", count))
	})

	server.POST("/decrement", func(rc *atreugo.RequestCtx) error {
		count--
		return rc.TextResponse(fmt.Sprintf("%d", count))
	})

	server.GET("/", func(rc *atreugo.RequestCtx) error {
		head := elem.Head(nil, elem.Script(elem.Attrs{attrs.Src: "https://unpkg.com/htmx.org@1.9.6"}))
		bodyStyle := elem.Style{
			styles.BackgroundColor: "#f4f4f4",
			styles.FontFamily:      "Arial, sans-serif",
			styles.Height:          "100vh",
			styles.Display:         "flex",
			styles.FlexDirection:   "column",
			styles.AlignItems:      "center",
			styles.JustifyContent:  "center",
		}

		buttonStyle := elem.Style{
			styles.Padding:         "10px 20px",
			styles.BackgroundColor: "#007BFF",
			styles.Color:           "#fff",
			styles.BorderColor:     "#007BFF",
			styles.BorderRadius:    "5px",
			styles.Margin:          "10px",
			styles.Cursor:          "pointer",
		}

		body := elem.Body(
			elem.Attrs{
				attrs.Style: elem.ApplyStyle(bodyStyle),
			},
			elem.H1(nil, elem.Text("Counter App")),
			elem.Div(elem.Attrs{attrs.ID: "count"}, elem.Text("0")),
			elem.Button(
				elem.Attrs{
					htmx.HXPost:   "/increment",
					htmx.HXTarget: "#count",
					attrs.Style:   elem.ApplyStyle(buttonStyle),
				},
				elem.Text("+"),
			),
			elem.Button(
				elem.Attrs{
					htmx.HXPost:   "/decrement",
					htmx.HXTarget: "#count",
					attrs.Style:   elem.ApplyStyle(buttonStyle),
				},
				elem.Text("-"),
			),
		)

		pageContent := elem.Html(nil, head, body)
		html := pageContent.Render()

		return rc.HTTPResponse(html)
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
