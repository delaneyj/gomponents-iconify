package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseTagBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16 6c0-1.886 0-2.828-.586-3.414C14.828 2 13.886 2 12 2c-1.886 0-2.828 0-3.414.586C8 3.172 8 4.114 8 6"/><path stroke-linejoin="round" d="M10 15H6c-.471 0-.707 0-.854.146C5 15.293 5 15.53 5 16v1c0 .471 0 .707.146.854C5.293 18 5.53 18 6 18h4c.471 0 .707 0 .854-.146C11 17.707 11 17.47 11 17v-1c0-.471 0-.707-.146-.854C10.707 15 10.47 15 10 15Z"/><path stroke-linecap="round" d="M6 6.5V15m0 6.5v-3m12-12V10m0 11.5V14M2 14c0-3.771 0-5.657 1.172-6.828C4.343 6 6.229 6 10 6h4c3.771 0 5.657 0 6.828 1.172c.654.653.943 1.528 1.07 2.828M22 14c0 3.771 0 5.657-1.172 6.828C19.657 22 17.771 22 14 22h-4c-3.771 0-5.657 0-6.828-1.172c-.654-.653-.943-1.528-1.07-2.828"/></g>`),
		g.Group(children),
	)
}