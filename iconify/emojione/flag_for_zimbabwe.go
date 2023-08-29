package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForZimbabwe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<defs><path id="emojioneFlagForZimbabwe0" d="M11.3 30.1c.7 1.5.9 1.6 2.6 2.2c1.7.6 6.6 1.2 7.1 1.2c-.3-.3-5.8-4.4-6.9-5.2c-1.1-.8-1.4-1.2-1.4-1.2s-1-1.8-1.4-2.4c-.4-.5-1-1.3-1.8-1.3s-1.4.6-1.7 1.5c-.1 0-.9.5-.8.9c.2 0 1.5-.4 2 1.1s.5 3.9 0 5.1c-.5 1.2-1.2 2.1-1.2 3.2c0 .1.1.4.3.4c.1.5.8 3.8.8 3.8h9.8l1.6-6"/></defs><path fill="#75a843" d="M32 2c-6.2 0-11.9 1.9-16.6 5h33.3C43.9 3.9 38.2 2 32 2zm0 60c6.2 0 11.9-1.9 16.6-5H15.4c4.7 3.1 10.4 5 16.6 5"/><path fill="#ed4c5c" d="M7 47h51c1.8-3.1 3-6.5 3.6-10.1H3.4C4 40.6 5.3 44 7 47z"/><path fill="#3e4347" d="M3 32c0 1.7.2 3.4.4 5h58.1c.3-1.6.4-3.3.4-5s-.2-3.4-.4-5H3.4c-.2 1.6-.4 3.3-.4 5"/><path fill="#ed4c5c" d="M3.4 27h58.2c-.6-3.6-1.9-7-3.6-10H7c-1.7 3-3 6.4-3.6 10"/><path fill="#ffce31" d="M12.3 17H58c-2.3-4-5.6-7.5-9.5-10h-33c-1.1.7-2.2 1.6-3.2 2.4V17m0 30v7.6c1 .9 2.1 1.7 3.2 2.4h33c3.9-2.6 7.1-6 9.5-10.1H12.3z"/><path fill="#f9f9f9" d="M12.6 12.2H9.5C4.8 17.5 2 24.4 2 32c0 7.7 2.9 14.8 7.8 20.1h2.5L32.4 32L12.6 12.2z"/><path fill="#3e4347" d="M12.3 9.4c-1 .9-2 1.8-2.8 2.8L29.3 32L9.5 51.8c.9 1 1.8 1.9 2.8 2.8L34.9 32L12.3 9.4z"/><path fill="#ed4c5c" d="m17.5 33.9l5.3-4.5l-6.3-1.1l-3.3-6.6l-3.3 6.6l-6.3 1.1L9 33.9l-1.4 6.7l5.6-3.2l5.7 3.2z"/><g fill="#ffce31"><use href="#emojioneFlagForZimbabwe0"/><g stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width=".2"><use href="#emojioneFlagForZimbabwe0"/><path d="M8.1 35.7s2.1-.1 2.7-.1c.6 0 5.8 0 8.8.7m-10.7-.7s-.2-.3-.1-.8s1.6-1.8 1.6-3.7m-.3 4.4s-.2-.8 1.7-1.5c.5-.2 2-1 2.1-1.7"/><path d="M11.6 34.1s4.6 0 8.3.8m-7.1-1.4s1.6.1 1.8 0c.3-.1.7-.3.9-.8"/><path d="M13.7 33.6s.9-.4 1-1m-3.5 6.9s-.4-2.5-.4-3.9l.9.9l1-.9l.9.9l1-.8l.9 1l1-.8l.8.9l1-.7l1.2 1.1"/><path d="M19.1 37.9c-3.5-.5-8.2-.5-8.2-.5m-1.6-8.7s2.9-1.4 3.4-1.6"/><path d="m10.8 36.5l1 .9l.8-.9l.9.9l1-.8l.8.9l1-.8l.8 1l1-.8l1 .9"/><circle cx="13.5" cy="34.9" r=".4"/><circle cx="15.2" cy="35" r=".4"/><circle cx="9.2" cy="24.7" r=".3"/></g></g>`),
		g.Group(children),
	)
}