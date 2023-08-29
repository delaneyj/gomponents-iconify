package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCollectionOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileCollectionOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 23v-9L31 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h12"/><path fill="#fff" d="M30.85 30C28.724 30 27 32.009 27 34.486c0 4.487 4.55 8.565 7 9.514c2.45-.949 7-5.027 7-9.514C41 32.01 39.276 30 37.15 30c-1.302 0-2.453.753-3.15 1.906C33.303 30.753 32.152 30 30.85 30Z"/><path d="M30 4v10h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileCollectionOne0)"/>`),
		g.Group(children),
	)
}