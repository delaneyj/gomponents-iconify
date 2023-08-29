package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilGiftCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M6 8h14a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1Zm14 12V9H6v11h14Z" clip-rule="evenodd"/><path d="M12.5 8.5h1v12h-1z"/><path d="M5 13.5h16v1H5z"/><path fill-rule="evenodd" d="M11.943 4.554c1.125.978 1.787 3.001 1.035 3.866c-.75.864-2.374.27-3.659-.847c-1.094-.951-1.469-2.399-.712-3.269c.756-.87 2.242-.701 3.336.25ZM9.975 6.82c.897.779 2.033 1.194 2.249.945c.3-.344-.174-1.792-.937-2.455c-.725-.63-1.595-.73-1.926-.349c-.33.38-.11 1.228.614 1.859Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14.298 4.554c-1.124.978-1.787 3.001-1.035 3.866c.75.864 2.374.27 3.66-.847c1.094-.951 1.469-2.399.712-3.269c-.756-.87-2.242-.701-3.337.25Zm1.969 2.265c-.897.779-2.033 1.194-2.249.945c-.3-.344.174-1.792.936-2.455c.725-.63 1.596-.73 1.926-.349c.33.38.111 1.228-.613 1.859Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilGiftCircleFilled0)"/></g>`),
		g.Group(children),
	)
}