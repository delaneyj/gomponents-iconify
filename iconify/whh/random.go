package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Random(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M766 501q-25 27-62-13V322q-123 10-197 52q-15-61-70-110q112-72 267-73V23q37-39 62-12l243 207q15 16 15 38t-15 38zm-318 75q0 46 72.5 81T704 701V535q37-39 62-12l243 207q15 16 15 38t-15 38l-243 207q-25 27-62-13V832q-104 0-192.5-34.5t-140-93.5T320 576V448q0-47-72.5-82T64 322v-2q-26 0-45-19T0 255.5T19 210t45-19q104 0 192.5 34.5t140 93.5T448 448v128zM64 704v-3q123-10 197-52q15 60 70 110q-112 73-267 73q-26 0-45-19T0 767.5t19-45T64 704z"/>`),
		g.Group(children),
	)
}