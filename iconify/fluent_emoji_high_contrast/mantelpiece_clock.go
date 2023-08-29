package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MantelpieceClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" d="M29.5 24.82v1.35c0 .184-.146.33-.33.33h-.653l-.002 1a.998.998 0 0 1 .995 1c0 .554-.446 1-1 1h-25c-.554 0-1-.446-1-1c0-.552.444-.997.995-1l-.002-1H2.83a.328.328 0 0 1-.33-.33v-1.34c0-.184.146-.33.33-.33a2.668 2.668 0 0 0 2.67-2.67V13.5h.01V13c0-.603.048-1.214.153-1.824c1.012-5.708 6.463-9.525 12.17-8.513c5.096.897 8.677 5.335 8.667 10.336l-.001.501h.011v8.33a2.668 2.668 0 0 0 2.67 2.67c.186 0 .324.146.32.32ZM16.698 12.7l.13.3l-.13.3a.75.75 0 1 1-.988-.988l.3-.13l.3.13a.754.754 0 0 1 .388.388Zm-.688 9.8a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19Z"/>`),
		g.Group(children),
	)
}