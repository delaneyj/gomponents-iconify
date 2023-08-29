package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modularplug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M128 226c-54.124 0-98-43.876-98-98s43.876-98 98-98s98 43.876 98 98s-43.876 98-98 98zm-1-8c49.706 0 90-40.294 90-90s-40.294-90-90-90s-90 40.294-90 90s40.294 90 90 90z"/><path d="M128 197c-38.108 0-69-30.668-69-68.5S89.892 60 128 60c38.108 0 69 30.668 69 68.5S166.108 197 128 197zm0-6.29c34.22 0 61.96-27.54 61.96-61.511c0-33.971-27.74-61.51-61.96-61.51s-61.96 27.539-61.96 61.51s27.74 61.51 61.96 61.51z"/><path d="M128 170c-23.196 0-42-18.804-42-42s18.804-42 42-42s42 18.804 42 42s-18.804 42-42 42z"/></g>`),
		g.Group(children),
	)
}