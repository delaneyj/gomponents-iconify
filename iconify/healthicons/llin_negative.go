package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LlinNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsLlinNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 15c1.223 0 2.343-.264 3.212-.7l.84 2.516a1 1 0 0 0 1.897-.632l-1-3a.997.997 0 0 0-.166-.307c.141-.277.217-.572.217-.877c0-1.451-1.718-2.662-4-2.94V6h-2v3.06c-2.282.278-4 1.489-4 2.94c0 .305.076.6.217.877a.997.997 0 0 0-.166.307l-1 3a1 1 0 0 0 1.898.632l.839-2.517c.869.438 1.989.701 3.212.701Zm-6.777 3.578a1 1 0 0 1 .077 1.412l-2.11 2.355a1 1 0 1 1-1.49-1.335l2.11-2.355a1 1 0 0 1 1.413-.077Zm-5.776 5.816a1 1 0 1 0-.894-1.788L8 23.882V27.5a1 1 0 1 0 2 0v-2.382l1.447-.724ZM6.106 41.947a1 1 0 0 1 .447-1.341L8 39.882V37.5a1 1 0 1 1 2 0v3.618l-2.553 1.276a1 1 0 0 1-1.341-.447ZM10 31a1 1 0 1 0-2 0v3a1 1 0 1 0 2 0v-3Zm20.7-11.01a1 1 0 0 1 1.49-1.335l2.11 2.355a1 1 0 0 1-1.49 1.335L30.7 19.99Zm5.406 3.063a1 1 0 0 0 .447 1.341l1.447.724V27.5a1 1 0 1 0 2 0v-3.618l-2.553-1.276a1 1 0 0 0-1.341.447Zm5.341 17.553a1 1 0 1 1-.894 1.788L38 41.118V37.5a1 1 0 1 1 2 0v2.382l1.447.724ZM39 30a1 1 0 0 0-1 1v3a1 1 0 1 0 2 0v-3a1 1 0 0 0-1-1Zm-23-3v15h2v-4h13v4h2v-6H18v-9h-2Zm4 5a1.5 1.5 0 0 0 0 3h12a1.5 1.5 0 0 0 0-3H20Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsLlinNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}