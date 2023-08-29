package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiBookmark0" d="m46.5 56l-10-11.151L26.5 56V10.958h20z"/></defs><path fill="#EA5A47" d="m46.5 56l-10-11.151L26.5 56V10.958h20z"/><path fill="#D22F27" d="M41.864 12.03v37.854l4.523 5.044V12.03z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiBookmark0"/><use href="#openmojiBookmark0"/></g>`),
		g.Group(children),
	)
}