package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#0576B4"/><g fill="#FFF"><path d="m24.668 18.863l-1.059-.663l2.31-.027l.048-.026v.025l.954-.01l-1.619 2.634l-.145-1.27l-7.364 4.063L15 19.86l-7.83 4.126v-.94l8.073-4.253l2.792 3.729l6.634-3.659zm-10.112-.905l-3.06 1.611V8.644h3.06v9.314zm8.653.481l-3.06 1.7V8.644h3.06v9.795z"/><path d="m18.883 20.843l-.657.364l-2.404-3.21V9.894h3.06v10.949zm-8.654-.607l-3.06 1.612V11.312h3.06v8.924z" opacity=".5"/></g></g>`),
		g.Group(children),
	)
}