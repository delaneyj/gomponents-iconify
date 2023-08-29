package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupMusicLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 16c-5.76 0-6.78-5.74-6.96-10.294c-.051-1.266-.076-1.9.4-2.485c.475-.586 1.044-.682 2.183-.874A26.374 26.374 0 0 1 12 2c1.784 0 3.253.157 4.377.347c1.139.192 1.708.288 2.184.874c.476.586.45 1.219.4 2.485c-.18 4.553-1.2 10.294-6.96 10.294Z"/><path stroke-linecap="round" d="M13 10.5V5"/><circle cx="11.5" cy="10.5" r="1.5"/><path stroke-linecap="round" d="M15 7a2 2 0 0 1-2-2"/><path d="m19 5l.949.316c.99.33 1.485.495 1.768.888c.283.393.283.915.283 1.958v.073c0 .86 0 1.291-.207 1.643c-.207.352-.584.561-1.336.98L17.5 12.5M5 5l-.949.316c-.99.33-1.485.495-1.768.888C2 6.597 2 7.12 2 8.162v.073c0 .86 0 1.291.207 1.643c.207.352.584.561 1.336.98L6.5 12.5" opacity=".5"/><path stroke-linecap="round" d="M12 16v3" opacity=".5"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.5 22h-7l.34-1.696a1 1 0 0 1 .98-.804h4.36a1 1 0 0 1 .98.804L15.5 22Z"/><path stroke-linecap="round" d="M18 22H6" opacity=".5"/></g>`),
		g.Group(children),
	)
}