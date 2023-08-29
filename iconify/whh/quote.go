package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M855.048 768h-87l-256 256V768h-343q-57 0-113-57t-56-115V172q0-58 56-115t113-57h686q57 0 113 57t56 115v424q0 58-56 115t-113 57zm-407-384q0-27-18.5-45.5t-45.5-18.5h-64v-64h32q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-96q-26 0-45 19t-19 45v256q0 26 19 45t45 19h128q26 0 45-19t19-45V384zm384 0q0-27-19-45.5t-45-18.5h-64v-64h32q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-96q-27 0-45.5 19t-18.5 45v256q0 26 18.5 45t45.5 19h128q26 0 45-18.5t19-45.5V384z"/>`),
		g.Group(children),
	)
}