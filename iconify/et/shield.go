package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13.5 0C4.563 0 .292 5.047.114 5.262A.506.506 0 0 0 0 5.581v10.054c0 8.504 7.826 13.553 11.19 15.329a.501.501 0 0 0 .467-.885C7.677 27.978 1 23.308 1 15.634V5.771c.702-.755 4.632-4.597 12-4.765V31.5a.5.5 0 0 0 .875.331C15.512 31.127 27 25.816 27 15.5v-10a.5.5 0 0 0-.115-.319C26.737 5.002 23.171.806 15.744.057a.506.506 0 0 0-.548.447a.5.5 0 0 0 .447.548c6.287.634 9.703 3.945 10.357 4.64V15.5c0 8.952-9.363 13.949-12 15.179V.5a.5.5 0 0 0-.5-.5z"/>`),
		g.Group(children),
	)
}