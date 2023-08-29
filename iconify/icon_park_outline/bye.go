package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="m35 26.614l-19.854-19.3a2.928 2.928 0 0 0-4.259.188a3.334 3.334 0 0 0 .18 4.544l10.024 9.93"/><path stroke-linejoin="round" d="M21.09 21.976L10.178 11.155a3.486 3.486 0 0 0-4.735-.161a3.032 3.032 0 0 0-.18 4.417l10.993 11.203"/><path d="M16.255 26.614L10 20.5a3.299 3.299 0 0 0-4.633-.08a3.233 3.233 0 0 0-.093 4.588c9.23 9.536 14.02 14.04 15.817 15.545C24 42.99 29.735 44 32.73 42c2.995-2 5.702-4.846 6.987-7.671c.765-1.683 2.25-6.87 4.458-15.561a3.305 3.305 0 0 0-2.46-4.034a3.5 3.5 0 0 0-4.166 2.493L35 26.614m-3.284-13.948a19.597 19.597 0 0 0-8.752-7.187M5.194 33.776a19.599 19.599 0 0 0 8.364 7.635"/></g>`),
		g.Group(children),
	)
}