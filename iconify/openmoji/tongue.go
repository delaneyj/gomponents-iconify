package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tongue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiTongue0" d="M62.417 12.5C56.28 19.625 48.969 22.19 36 22.188c-12.969.001-20.281-2.563-26.417-9.688"/></defs><path fill="none" d="M15.625 21.75h41.75"/><path fill="#ea5a47" d="M56.967 16.777c0 6-41.78 7.82-41.342 1.007C5.97 53.023 24.962 60.375 36.5 60.375s31.895-6.02 20.467-43.598z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiTongue0"/><use href="#openmojiTongue0"/><path d="M57.375 21.75C66.5 55.188 48.038 60.375 36.5 60.375s-30-5.188-20.875-38.625M36.5 60.375V34.196"/></g>`),
		g.Group(children),
	)
}