package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wakeonlan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 45.491a21.524 21.524 0 0 1-21.5-21.5c0-8.035 5.175-13.835 9.067-17.5a.774.774 0 0 1 .49-.186a.687.687 0 0 1 .479.215l3.52 3.743a.685.685 0 0 1-.028.969c-4.854 4.565-7.018 8.5-7.018 12.758a14.99 14.99 0 1 0 29.98 0c0-5.589-3.752-9.965-6.903-12.654a.684.684 0 0 1-.075-.966l3.34-3.908a.687.687 0 0 1 .965-.077c7.59 6.48 9.183 13.15 9.183 17.604A21.525 21.525 0 0 1 24 45.49Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.604 2.509a.685.685 0 0 1 .685.685v.514h.822a.686.686 0 0 1 .686.685v11.992a.686.686 0 0 1-.686.685h-.257v9.524a.685.685 0 0 1-.686.685H20.83a.687.687 0 0 1-.685-.68l-.081-9.7h-.176a.685.685 0 0 1-.685-.685V4.564a.685.685 0 0 1 .685-.685h.77v-.685a.685.685 0 0 1 .686-.685Zm-7.401 7.242h9.594"/>`),
		g.Group(children),
	)
}