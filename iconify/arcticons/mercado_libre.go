package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MercadoLibre(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="19.5" ry="12.978"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.704 15.53a20.834 20.834 0 0 0 6.386 1.866a22.82 22.82 0 0 0 4.546-.773"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.882 15.614a8.616 8.616 0 0 1-5.165 1.485c-3.335 0-6.225-2.199-9.215-2.199c-2.668 0-7.189 4.373-7.189 5.164s1.31 1.26 2.372.74c.621-.303 3.31-2.914 5.484-2.914s9.219 7.136 9.857 7.806c.989 1.038-.926 3.274-2.149 2.05s-3.41-3.162-3.41-3.162"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.4 22.683a23.998 23.998 0 0 0-8.547 2.692m-2.273 2.081c.989 1.037-.926 3.273-2.149 2.05s-2.581-2.513-2.581-2.513"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.135 29.215c.988 1.037-.927 3.273-2.15 2.05s-2.025-1.962-2.025-1.962"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.201 31.316a2.309 2.309 0 0 0 3.649-.186"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.201 31.316c.531-.697.49-3.182-2.243-2.688c.642-1.219.066-3.146-2.388-2.01a1.69 1.69 0 0 0-3.146-.658a1.454 1.454 0 0 0-2.8-.28c-.544 1.103.296 3.096 2.092 1.976c-.182 1.944.84 2.537 2.684 1.78c.099 1.91 1.367 1.745 2.273 1.3a1.938 1.938 0 0 0 3.529.58Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.67 22.279a18.308 18.308 0 0 1 9.064 3.214"/>`),
		g.Group(children),
	)
}