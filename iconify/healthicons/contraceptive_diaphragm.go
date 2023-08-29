package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContraceptiveDiaphragm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M38 24.043c0 .38-.015.742-.045 1.087l.01.002c.023.358.035.72.035 1.084c0 1.409-.176 2.547-.483 3.468a1 1 0 0 0 1.898.632c.39-1.17.585-2.53.585-4.1c0-.161-.002-.323-.006-.483C44.406 28.11 47.176 37 24.27 37C1.429 37 3.199 27.669 8.007 25.711a18.76 18.76 0 0 0-.007.505c0 1.31.136 2.47.406 3.498a1 1 0 0 0 1.934-.509c-.218-.83-.34-1.815-.34-2.989c0-.342.01-.681.03-1.017l.02-.004a12.77 12.77 0 0 1-.05-1.152C10 16.287 16.268 10 24 10s14 6.287 14 14.043Zm-17.374-6.767a1 1 0 1 0-.576-1.915A8.5 8.5 0 0 0 14 23.428a1 1 0 0 0 2 .017a6.5 6.5 0 0 1 4.626-6.17Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}