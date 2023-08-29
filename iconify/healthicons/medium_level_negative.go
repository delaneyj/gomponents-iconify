package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediumLevelNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsMediumLevelNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM11.065 13.48A17.926 17.926 0 0 0 6.109 24h3.944v2H6a17.902 17.902 0 0 0 3.864 11.127l2.211-2.21l.708-.708l1.414 1.414l-.707.707l-2.233 2.233l-.066.066l.02.019l-1.421 1.407A19.923 19.923 0 0 1 4 25.99C4 14.95 12.955 6 24 6s20 8.949 20 19.99a19.928 19.928 0 0 1-6.085 14.357l-1.391-1.437l.013-.012L36 38.36l-2.233-2.233l-.708-.707l1.415-1.414l.707.707l2.233 2.233l.48.48A17.905 17.905 0 0 0 42 26h-4.316v-2h4.207a17.926 17.926 0 0 0-4.956-10.52l-2.228 2.227l-.707.707L32.586 15l.707-.707l2.172-2.172A17.924 17.924 0 0 0 25 8.027V13h-2V8.027a17.925 17.925 0 0 0-10.465 4.094l.172.172l2.121 2.121l.708.707l-1.415 1.414l-.707-.707l-2.121-2.12l-.228-.229Zm9.561 24.224c-.028.392.088.781.326 1.094l1.922 2.212c1.126.49 1.626.49 2.569 0l1.922-2.212c.238-.313.354-.702.326-1.094L25.31 23.902c-.145-1.997-2.17-2.015-2.314-.018l-2.37 13.82Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsMediumLevelNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}