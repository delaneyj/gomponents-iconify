package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Synagogue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiSynagogue0" stroke-linecap="round" stroke-linejoin="round" d="m35 47l8.66-15H26.34z"/></defs><g stroke-miterlimit="10" stroke-width="2"><path fill="#F4AA41" d="M52 33.78V60H18V33.78a1 1 0 0 1 .408-.805L34.38 21.23a.946.946 0 0 1 1.238 0l15.973 11.744a1 1 0 0 1 .408.806zM8 39h10v21H8zm44 0h10v21H52z"/><path fill="#A57939" d="M39 60h-8v-6a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v6z"/><path fill="#3F3F3F" d="m57 30l-5 9h10zm-44 0l-5 9h10z"/><path fill="#FFF" d="m35 27l-8.66 15h17.32z"/><path fill="#D0CFCE" d="m35 47l8.66-15H26.34z"/><path fill="#D0CFCE" d="m35 27l-8.66 15h17.32z"/></g><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><path d="M52 33.78V60H18V33.78a1 1 0 0 1 .408-.805L34.38 21.23a.946.946 0 0 1 1.238 0l15.973 11.744a1 1 0 0 1 .408.806z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 39h10v21H8zm44 0h10v21H52zM39 60h-8v-6a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v6z"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 47h-2v-3a1 1 0 0 1 1-1h0m45 4h-2v-3a1 1 0 0 1 1-1h0M35 60v-9m22-21l-5 9h10zm-44 0l-5 9h10zm22-3l-8.66 15h17.32z"/><use href="#openmojiSynagogue0" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="m35 27l-8.66 15h17.32z"/><use href="#openmojiSynagogue0" stroke-linecap="round" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}