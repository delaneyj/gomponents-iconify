package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UiPreferencesNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsUiPreferencesNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm16 32a4 4 0 1 1-5-3.874V7a1 1 0 1 1 2 0v21.126c1.725.444 3 2.01 3 3.874Zm7-9.078c0-.526.474-.922 1-.922s1 .396 1 .922V41a1 1 0 1 1-2 0V22.922ZM28 16a4.002 4.002 0 0 0-3-3.874V7a1 1 0 1 0-2 0v5.126A4.002 4.002 0 0 0 24 20a4 4 0 0 0 4-4Zm8 17c-.526 0-1 .396-1 .922V41a1 1 0 1 0 2 0v-7.078c0-.526-.474-.922-1-.922Zm1-9.874A4.002 4.002 0 0 1 36 31a4 4 0 0 1-1-7.874V7a1 1 0 1 1 2 0v16.126ZM11 38.922c0-.526.474-.922 1-.922s1 .396 1 .922V41a1 1 0 1 1-2 0v-2.078Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsUiPreferencesNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}