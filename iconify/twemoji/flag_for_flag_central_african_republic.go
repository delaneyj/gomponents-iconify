package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFlagCentralAfricanRepublic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#003082" d="M32 5H4a4 4 0 0 0-4 4v2.5h36V9a4 4 0 0 0-4-4z"/><path fill="#EEE" d="M0 11.5h36V18H0z"/><path fill="#289728" d="M0 18h36v6.5H0z"/><path fill="#FFCE00" d="M0 24.5V27a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-2.5H0z"/><path fill="#D21034" d="M15 5h6v26h-6z"/><path fill="#FFCE00" d="M6.878 7.612l-.68-2.094l-.681 2.094H3.316l1.781 1.294L4.417 11l1.781-1.294L7.979 11l-.681-2.094L9.08 7.612z"/>`),
		g.Group(children),
	)
}