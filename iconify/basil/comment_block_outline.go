package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentBlockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.534 4.131c.098-.143-.01-.343-.182-.322c-5.395.634-8.879 6.295-6.76 11.495l.932 2.289a.25.25 0 0 1-.075.289l-1.971 1.583a.75.75 0 0 0 .47 1.335h7.82c4.798 0 8.718-3.84 8.97-8.58c.009-.179-.207-.271-.343-.156c-.317.271-.663.51-1.032.71a.43.43 0 0 0-.219.306a7.484 7.484 0 0 1-7.376 6.22H6.079l.31-.248a1.75 1.75 0 0 0 .524-2.025l-.932-2.289c-1.623-3.983.757-8.296 4.68-9.28a.433.433 0 0 0 .294-.254c.157-.379.352-.738.58-1.073Z"/><path fill="currentColor" fill-rule="evenodd" d="M12.832 10.107a4.5 4.5 0 1 1 7.336-5.215a4.5 4.5 0 0 1-7.336 5.215Zm2.144-.022a3 3 0 0 0 4.109-4.109l-4.109 4.109Zm3.048-5.17l-4.109 4.109a3 3 0 0 1 4.109-4.109Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}