package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snikket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.75 42.624a21.488 21.488 0 1 1 8.072-8.224"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.75 42.624a21.507 21.507 0 0 1 0-37.248"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.011 23.3a21.494 21.494 0 0 1 18.81 11.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.286 28.733a21.515 21.515 0 0 1-24.263 2.982"/><circle cx="17.243" cy="13.617" r="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.804 13.617a1.556 1.556 0 0 1-1.264-2.47a2.467 2.467 0 0 0-.297-.03a2.5 2.5 0 1 0 2.5 2.5a2.467 2.467 0 0 0-.03-.297a1.548 1.548 0 0 1-.91.297ZM4.481 33.041a11.932 11.932 0 0 1 1.017 4.839a11.941 11.941 0 0 1-2.633 7.49a12.025 12.025 0 0 0 1.671.13a11.958 11.958 0 0 0 8.165-3.205"/>`),
		g.Group(children),
	)
}