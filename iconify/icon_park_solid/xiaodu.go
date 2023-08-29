package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xiaodu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSXiaodu0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4.143 21.08c.355-3.786.533-5.68 2.513-7.705c1.98-2.025 4.16-2.28 8.517-2.788C18.003 10.257 21.131 10 24 10c2.87 0 5.996.257 8.827.587c4.358.508 6.537.763 8.517 2.788c1.98 2.026 2.158 3.919 2.513 7.706c.09.955.143 1.937.143 2.919s-.054 1.964-.143 2.92c-.355 3.786-.533 5.68-2.513 7.705c-1.98 2.025-4.16 2.28-8.517 2.788c-2.83.33-5.957.587-8.827.587s-5.996-.257-8.827-.587c-4.358-.508-6.537-.763-8.517-2.788c-1.98-2.026-2.158-3.919-2.513-7.706A31.388 31.388 0 0 1 4 24c0-.982.054-1.964.143-2.92Z"/><path stroke="#000" d="M16 19v10m17-10l-5 5l5 5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSXiaodu0)"/>`),
		g.Group(children),
	)
}