package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeGrinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiCoffeeGrinder0" d="m38 54.5l-.229-.973A1 1 0 0 0 37 54.5h1Zm17-4h1a1 1 0 0 0-1.229-.973L55 50.5ZM39 58v-3.5h-2V58h2Zm1 1a1 1 0 0 1-1-1h-2a3 3 0 0 0 3 3v-2Zm13 0H40v2h13v-2Zm1-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3h-2Zm0-3.5V58h2v-3.5h-2Zm0-4v4h2v-4h-2Zm-15.771 4.973l17-4l-.458-1.946l-17 4l.458 1.946Z"/></defs><path fill="#fff" d="M38 54V41h17v9l-17 4Z"/><path fill="#A57939" d="M38 19h17v23H38V19Z"/><g fill="#6A462F"><path fill-rule="evenodd" d="M38 54.5V58a2 2 0 0 0 2 2h13a2 2 0 0 0 2-2v-7.5l-17 4Z" clip-rule="evenodd"/><use href="#openmojiCoffeeGrinder0"/><use href="#openmojiCoffeeGrinder0"/></g><path fill="#A57939" d="M22 13.528C22 16.556 19.761 22 17 22s-5-5.444-5-8.472c0-3.027 2.239-2.491 5-2.491s5-.536 5 2.491Z"/><path fill="#d0cfce" d="M44 14h5v5h-5z"/><g stroke="#000" stroke-width="2"><path fill="none" stroke-linecap="round" stroke-linejoin="round" d="m41 53.794l11-2.588"/><path fill="none" d="M38 21a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2v37a2 2 0 0 1-2 2H40a2 2 0 0 1-2-2V21Z"/><path fill="none" stroke-linejoin="round" d="M44 14h5v5h-5zm0 2h-6l-15 9h-6v-3"/><path fill="none" d="M22 13.528C22 16.556 19.761 22 17 22s-5-5.444-5-8.472c0-3.027 2.239-2.491 5-2.491s5-.536 5 2.491Z"/><path stroke-linecap="round" d="M39 41h13"/></g>`),
		g.Group(children),
	)
}