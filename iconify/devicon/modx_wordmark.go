package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModxWordmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#00b5de" d="m92.1 40.8l18.3-29.3h-44l-6.1 9.8z" class="modx-original-wordmark-st0"/><path fill="#ff5529" d="M101.8 93.8v-44l-9.5-5.9l-20 31.5z" class="modx-original-wordmark-st2"/><path fill="#00decc" d="M30.9 2.7v44l10.2 5.9l51-11.8z" class="modx-original-wordmark-st3"/><path fill="#ff9640" d="M40.9 56L22.6 84.8h44l25.7-40.9z" class="modx-original-wordmark-st4"/><path fill="#102c53" d="M16.2 122.1h2.6l1.3-15.6h-2.6v-3h7l4.8 12.2c.4 1.1.5 1.8.5 1.8h.1s.2-.7.5-1.8l4.8-12.2h7v3h-2.6l1.2 15.6h2.6v3h-8.9v-3h2.4l-.7-11c-.1-.8.1-2 .1-2h-.1s-.2 1.2-.5 1.9l-4.1 10h-3.4l-4.1-10c-.3-.7-.5-1.9-.5-1.9h-.1s.2 1.2.1 2l-.8 10.9h2.4v3h-9v-2.9zM56.1 103c6.4 0 11.2 4.8 11.2 11c0 6.4-4.8 11.3-11.2 11.3c-6.4 0-11.2-5-11.2-11.3c0-6.1 4.8-11 11.2-11zm0 18.9c3.9 0 7.1-3.4 7.1-7.8c0-4.3-3.2-7.5-7.1-7.5c-3.9 0-7.1 3.3-7.1 7.5c0 4.4 3.2 7.8 7.1 7.8zm13.2.2h2.6v-15.6h-2.6v-3h9.4c1.7 0 3.1.1 4.5.5c4.2 1.2 7 5 7 10.3c0 5.1-2.6 8.8-6.7 10.2c-1.5.5-3 .6-4.8.6h-9.4zm9.2-.3c1.4 0 2.5-.1 3.5-.6c2.5-1 4-3.5 4-7.1c0-3.7-1.6-6.2-4.2-7.1c-1.1-.4-2.1-.5-3.4-.5h-2.6v15.2zm14 .3h1.9l5.5-8l-5.3-7.6h-1.9v-3h8.2v3h-2.1l3 4.5c.4.5.6.9.6.9h.1s.2-.4.6-.9l2.9-4.5h-2.1v-3h7.9v3h-1.9l-5.5 7.9l5.4 7.7h1.9v3h-8.2v-3h2.1l-3-4.5c-.4-.5-.6-.9-.6-.9h-.1c-.1 0-.2.4-.6.9l-2.9 4.5h2.1v3h-7.9z"/>`),
		g.Group(children),
	)
}