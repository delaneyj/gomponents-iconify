package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildFriendlyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 22q-.65 0-1.075-.425T7 20.5q0-.65.425-1.075T8.5 19q.65 0 1.075.425T10 20.5q0 .65-.425 1.075T8.5 22Zm11 0q-.65 0-1.075-.425T18 20.5q0-.65.425-1.075T19.5 19q.65 0 1.075.425T21 20.5q0 .65-.425 1.075T19.5 22Zm-6-14.4L9.6 2.9q.9-.45 1.95-.675T13.8 2q.875 0 1.688.163t1.537.487q.55.25.588.775t-.438.95L13.5 7.6ZM12 18q-2.075 0-3.538-1.463T7 13V5.3l-.5-.6q-.35-.425-.588-.562T5.4 4q-.425 0-.75.238t-.5.587q-.15.325-.363.5T3.3 5.5q-.575 0-.925-.4t-.15-.875q.375-.975 1.225-1.6T5.4 2q.75 0 1.375.35t1.275 1.1l10.7 12.9q.425.5.163 1.075T18 18h-6Zm0-2h3.85L9 7.7V13q0 1.25.875 2.125T12 16Zm.425-4.15Z"/>`),
		g.Group(children),
	)
}