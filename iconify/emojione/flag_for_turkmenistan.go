package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTurkmenistan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#83bf4f" d="M2 32c0 6.1 1.8 11.8 5 16.6V15.4C3.8 20.2 2 25.9 2 32zM32 2c-5.5 0-10.6 1.5-15 4v52c4.4 2.6 9.5 4 15 4c16.6 0 30-13.4 30-30S48.6 2 32 2z"/><path fill="#ed4c5c" d="M7 15.4v33.2c2.6 3.9 6 7.1 10 9.4V6c-4 2.3-7.4 5.6-10 9.4"/><path fill="#fff" d="M37.2 8.8c2.5 2.1 4 5.3 3.8 8.8c-.3 6-5.4 10.7-11.4 10.5c-1.4-.1-2.7-.4-3.8-.9c1.9 1.6 4.4 2.6 7.1 2.6c6.1 0 11-4.9 11-11c-.1-4.4-2.8-8.3-6.7-10m-10.5 8.1l.5-1.5h1.5l-1.2-.9l.5-1.4l-1.3.9l-1.2-.9l.5 1.4l-1.3.9h1.5zm-.5-7.1l.5 1.5l.5-1.5h1.5l-1.2-.9l.5-1.5l-1.3.9l-1.2-.9l.5 1.5l-1.3.9zm5.6 2.9l.5 1.5l.5-1.5h1.5l-1.2-.9l.4-1.5l-1.2.9l-1.3-.9l.5 1.5l-1.2.9zm-9.9 4.6l.5-1.5l-1.3.9l-1.2-.9l.5 1.5l-1.3.9h1.6l.4 1.5l.5-1.5h1.6zm4.8 5.2l.5-1.4h1.5l-1.2-.9l.5-1.5l-1.3.9l-1.2-.9l.5 1.5l-1.3.9h1.5z"/><path fill="#ffe62e" d="M7 15.4v6l3-3zm10 6v-6l-3 3zM7 29.9v5.9l3-2.9zm10 5.9v-5.9l-3 3zM7 42.6v6l3-3zm10 6v-6l-3 3z"/><path fill="#83bf4f" d="m9.8 42.6l-2.3-3.4l2.3-3.3h4.4l2.3 3.3l-2.3 3.4z"/><path fill="#ffe62e" d="m9.8 28.8l-2.3-3.3l2.3-3.3h4.4l2.3 3.3l-2.3 3.3z"/><path fill="#ed4c5c" d="m10.4 27.9l-1.7-2.4l1.7-2.4h3.2l1.7 2.4l-1.7 2.4z"/><path fill="#fff" d="m11.3 26.5l-.6-1l.6-1h1.4l.6 1l-.6 1z"/><path fill="#83bf4f" d="m9.2 14.6l-.7-1.1L14 8.3l2 3l-2.2 3.3z"/><path fill="#ed4c5c" d="m9.9 13.7l-.8-1.1l4-3.8l1.7 2.5l-1.7 2.4z"/><path fill="#ffe62e" d="m9 49l-.7 1.1l5.4 5.2l2-2.9l-2.2-3.3z"/><path fill="#ed4c5c" d="M9.6 49.9L8.9 51l3.9 3.8l1.7-2.4l-1.6-2.4zm.8-8.3l-1.7-2.4l1.7-2.4h3.2l1.7 2.4l-1.7 2.4z"/><path fill="#ffe62e" d="m11.3 40.2l-.7-1l.7-1h1.3l.7 1l-.7 1zm-.4-28.3l-.4-.6L12 9.9l.8 1l-.6 1z"/><path fill="#fff" d="m10.4 51.7l-.4.6l1.6 1.5l.8-1.1l-.7-1z"/>`),
		g.Group(children),
	)
}