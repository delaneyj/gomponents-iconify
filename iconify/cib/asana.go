package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.042 16.87a6.962 6.962 0 1 0 0 13.922a6.96 6.96 0 0 0 0-13.922zm-18.084 0a6.962 6.962 0 0 0 0 13.922a6.967 6.967 0 0 0 6.964-6.964a6.962 6.962 0 0 0-6.964-6.958zm16-8.698a6.958 6.958 0 1 1-13.916 0c0-3.844 3.115-6.964 6.958-6.964s6.958 3.12 6.958 6.964z"/>`),
		g.Group(children),
	)
}