package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupFirstBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m11 8l1.5-1.5v4"/><path d="m19 5l.949.316c.99.33 1.485.495 1.768.888c.283.393.283.915.283 1.958v.073c0 .86 0 1.291-.207 1.643c-.207.352-.584.561-1.336.98L17.5 12.5M5 5l-.949.316c-.99.33-1.485.495-1.768.888C2 6.597 2 7.12 2 8.162v.073c0 .86 0 1.291.207 1.643c.207.352.584.561 1.336.98L6.5 12.5"/><path stroke-linecap="round" d="M12 16v3"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.5 22h-7l.34-1.696a1 1 0 0 1 .98-.804h4.36a1 1 0 0 1 .98.804L15.5 22Z"/><path stroke-linecap="round" d="M18 22H6M17 2.456c.741.141 1.181.297 1.56.765c.477.586.452 1.219.401 2.485c-.18 4.553-1.2 10.294-6.96 10.294S5.22 10.26 5.038 5.706c-.05-1.266-.075-1.9.4-2.485c.476-.586 1.045-.682 2.184-.874A26.374 26.374 0 0 1 12 2c.718 0 1.386.025 2 .068"/></g>`),
		g.Group(children),
	)
}