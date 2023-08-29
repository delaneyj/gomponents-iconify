package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M23 11a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V11Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M6 8h14a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1Zm14 12V9H6v11h14Z" clip-rule="evenodd"/><path d="M12.5 8.5h1v12h-1z"/><path d="M5 13.5h16v1H5z"/><path fill-rule="evenodd" d="M11.943 4.554c1.125.978 1.787 3.001 1.035 3.866c-.75.864-2.374.27-3.659-.847c-1.094-.951-1.469-2.399-.712-3.269c.756-.87 2.242-.701 3.336.25ZM9.975 6.82c.897.779 2.033 1.194 2.249.945c.3-.344-.174-1.792-.937-2.455c-.725-.63-1.595-.73-1.926-.349c-.33.38-.11 1.228.614 1.859Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14.298 4.554c-1.124.978-1.787 3.001-1.035 3.866c.75.864 2.374.27 3.66-.847c1.094-.951 1.469-2.399.712-3.269c-.756-.87-2.242-.701-3.337.25Zm1.969 2.265c-.897.779-2.033 1.194-2.249.945c-.3-.344.174-1.792.936-2.455c.725-.63 1.596-.73 1.926-.349c.33.38.111 1.228-.613 1.859Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}