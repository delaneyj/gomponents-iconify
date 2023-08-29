package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldLockedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.9q-.175 0-.325-.025t-.3-.075q-3.425-1.125-5.4-4.1T4 11.1V6.375q0-.625.363-1.125t.937-.725l6-2.25q.35-.125.7-.125t.7.125l6 2.25q.575.225.938.725T20 6.375V11.1q0 .25-.013.5t-.037.5q-.25-.05-.475-.075T19 12q-2.075 0-3.538 1.463T14 17v4.25q-.325.175-.675.3l-.7.25q-.15.05-.3.075T12 21.9Zm4.85.1q-.35 0-.6-.25t-.25-.6v-3.3q0-.35.25-.6t.6-.25H17v-1q0-.825.588-1.413T19 14q.825 0 1.413.588T21 16v1h.15q.35 0 .6.25t.25.6v3.3q0 .35-.25.6t-.6.25h-4.3ZM18 17h2v-1q0-.425-.288-.712T19 15q-.425 0-.713.288T18 16v1Z"/>`),
		g.Group(children),
	)
}