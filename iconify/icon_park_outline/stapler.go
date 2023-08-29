package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stapler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M42 35v8H6v-8h25m7.384-13.194L8.314 10.862l-2.736 7.517l24.175 8.8M8.37 10.778l14.512 5.38a.03.03 0 0 0 .04-.018v0a5.285 5.285 0 0 0-3.119-6.792l-9.502-3.523a.09.09 0 0 0-.115.053l-1.816 4.9Zm8.385 11.942l-1.368 3.759a1 1 0 0 1-1.282.597l-6.578-2.394a1 1 0 0 1-.598-1.281l1.368-3.76"/><circle cx="36" cy="30" r="7"/></g>`),
		g.Group(children),
	)
}