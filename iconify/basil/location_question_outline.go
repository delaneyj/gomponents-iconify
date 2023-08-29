package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationQuestionOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.524 8.857a8.288 8.288 0 0 1 8.26-7.607h.432a8.288 8.288 0 0 1 8.26 7.607a8.944 8.944 0 0 1-1.99 6.396l-4.793 5.861a2.187 2.187 0 0 1-3.386 0l-4.793-5.861a8.943 8.943 0 0 1-1.99-6.396Zm8.26-6.107A6.788 6.788 0 0 0 5.02 8.98a7.443 7.443 0 0 0 1.656 5.323l4.793 5.862a.687.687 0 0 0 1.064 0l4.793-5.862A7.443 7.443 0 0 0 18.98 8.98a6.788 6.788 0 0 0-6.765-6.23h-.432Z" clip-rule="evenodd"/><path fill="currentColor" d="M13 16a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M12 7.75c-.947 0-1.75.795-1.75 1.821a.75.75 0 1 1-1.5 0c0-1.814 1.435-3.321 3.25-3.321s3.25 1.507 3.25 3.321c0 1.431-.888 2.664-2.152 3.127a.735.735 0 0 0-.287.184c-.057.063-.061.102-.061.118a.75.75 0 0 1-1.5 0c0-.924.743-1.494 1.332-1.71a1.82 1.82 0 0 0 1.168-1.719c0-1.026-.803-1.821-1.75-1.821Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}