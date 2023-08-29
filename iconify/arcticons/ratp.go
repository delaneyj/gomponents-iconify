package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ratp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.301 13.142s3.344-1.776 8.048-.417c0 0 2.299.627 4.912 2.821c0 0 .837.732 1.881 1.254c0 0 .942.628 2.51.419c1.88-.314 2.927-.314 2.927-.314s1.357-.21 1.357 1.045c0 0 .21 1.15.314 1.987c0 0 0 .208.314.312l.94.314s.733.21.628.94l-.105 1.151s.313-.627.94-.523c0 0 .732.106.732.837c0 0 0 .835.21 1.045c0 0 .209.419 1.463.523a2.6 2.6 0 0 1 .94.418c.942.73.314 1.88.314 1.88s-1.776 3.868-2.613 7.527c0 0-1.567 5.853-.105 10.139"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.847 38.046c-.459.039-.923.058-1.392.058c-8.972 0-16.244-7.298-16.244-16.302c.158-2.511.61-4.81 1.502-6.836m2.49-3.87C16.182 7.666 20.567 5.5 25.456 5.5c8.972 0 16.244 7.298 16.244 16.302a16.306 16.306 0 0 1-10.47 15.243"/>`),
		g.Group(children),
	)
}