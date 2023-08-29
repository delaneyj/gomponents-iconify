package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermScanWifiOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20.175q-.375 0-.75-.15t-.675-.45l-9.15-9.15q-.6-.6-.587-1.438t.637-1.337q2.25-1.8 4.95-2.725T12 4q2.875 0 5.575.925t4.95 2.725q.625.5.638 1.338t-.588 1.437l-9.15 9.15q-.3.3-.675.45t-.75.15Zm0-2.025l9.1-9.1q-1.975-1.5-4.3-2.275T12 6q-2.475 0-4.8.775T2.9 9.05l9.1 9.1ZM12 15q.425 0 .713-.287T13 14v-3q0-.425-.288-.713T12 10q-.425 0-.713.288T11 11v3q0 .425.288.713T12 15Zm0-6q.425 0 .713-.287T13 8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8q0 .425.288.713T12 9Zm0 9.15Z"/>`),
		g.Group(children),
	)
}