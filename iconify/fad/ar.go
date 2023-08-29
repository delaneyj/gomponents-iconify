package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M25 184c-.47 2.68-.227 4.354 1 6s4.563 2.464 8 2c3.437-.464 5.078-.958 6-4c.922-3.042 40.911-105.217 40.911-105.217c.601-1.537 2.13-1.93 3.404-.886c0 0 134.305 110.087 136.602 110.103c1.194.008 3.15-.157 4.942-1.477c1.656-1.22 3.173-3.57 3.824-4.25c1.355-1.414 3.54-3.424.46-6.697C227.063 176.304 104.148 78.758 100 74c-4.148-4.758-10.631-10-16-10s-13.914 2.695-16 8c-2.086 5.305-42.53 109.32-43 112z"/>`),
		g.Group(children),
	)
}