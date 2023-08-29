package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 5h14a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1Zm14 12V6H3v11h14Z" clip-rule="evenodd"/><path d="M9.5 5.5h1v12h-1z"/><path d="M2 10.5h16v1H2z"/><path fill-rule="evenodd" d="M8.943 1.554c1.125.978 1.787 3.001 1.035 3.866c-.75.864-2.374.27-3.659-.847c-1.094-.951-1.469-2.399-.712-3.269c.756-.87 2.242-.701 3.336.25ZM6.975 3.82c.897.779 2.033 1.194 2.249.945c.3-.344-.174-1.792-.937-2.455c-.725-.63-1.595-.73-1.926-.349c-.33.38-.11 1.228.614 1.859Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M11.298 1.554c-1.124.978-1.787 3.001-1.035 3.866c.75.864 2.374.27 3.66-.847c1.094-.951 1.469-2.399.712-3.269c-.756-.87-2.242-.701-3.337.25Zm1.969 2.265c-.897.779-2.033 1.194-2.249.945c-.3-.344.174-1.792.936-2.455c.725-.63 1.596-.73 1.926-.349c.33.38.111 1.228-.613 1.859Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}