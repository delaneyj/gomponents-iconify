package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUpCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopThumbsUpCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M6.5 11a1 1 0 0 0-1 1v5.5a1 1 0 0 0 1 1H9a1 1 0 0 0 1-1V12a1 1 0 0 0-1-1H6.5Zm1.25 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm6.731-9.335a.167.167 0 0 0-.273.056l-.194 1.165A5.161 5.161 0 0 1 12.23 12v4.681l4.196.84a.754.754 0 0 1 .082.02a1.432 1.432 0 0 0 1.603-.63l.217-.353l.507-1.338a3.899 3.899 0 0 0-.018-2.809l-.31-.621a1.427 1.427 0 0 0-1.272-.79H15.73c-.043 0-.084-.003-.125-.008a1 1 0 0 1-1.096-1.235l.15-.594a1.98 1.98 0 0 0 .057-.547l-.014-.41a.849.849 0 0 0-.202-.521l-.02-.02Zm-3.428 10.82l4.947.989a3.432 3.432 0 0 0 3.815-1.516l.257-.416a1.003 1.003 0 0 0 .1-.206l.534-1.409a5.899 5.899 0 0 0-.044-4.292a.998.998 0 0 0-.037-.082l-.329-.659A3.427 3.427 0 0 0 17.231 9h-.525a4.02 4.02 0 0 0 .01-.453l-.014-.41a2.849 2.849 0 0 0-.764-1.844l-.043-.042c-1.19-1.192-3.226-.628-3.634 1.006a1.01 1.01 0 0 0-.017.079l-.203 1.221a3.161 3.161 0 0 1-1.354 2.104a1 1 0 0 0-.456.839v5.98a1.002 1.002 0 0 0 .822 1.004Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopThumbsUpCircleFilled0)"/></g>`),
		g.Group(children),
	)
}