package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UiUserProfileNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsUiUserProfileNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm11 35.63c0 .34.057.675.166.991A17.942 17.942 0 0 1 6 24c0-9.941 8.059-18 18-18s18 8.059 18 18a17.94 17.94 0 0 1-5.181 12.636a2.99 2.99 0 0 0-2.498-4.002c-7.758-.803-12.836-.88-20.632-.018A3.028 3.028 0 0 0 11 35.631ZM4 24c0 10.772 8.517 19.556 19.184 19.984a10.295 10.295 0 0 0 .678.015L24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24Zm20 4a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsUiUserProfileNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}