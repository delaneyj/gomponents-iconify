package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gigalife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.776 12.766v8.604c0 1.577-1.305 2.868-2.9 2.868c-.869 0-1.594-.287-2.029-.86"/><path d="M14.877 12.766c1.594 0 2.899 1.29 2.899 2.868v1.864c0 1.577-1.305 2.868-2.9 2.868s-2.898-1.29-2.898-2.868v-1.864c0-1.577 1.304-2.868 2.899-2.868Z"/></g><circle cx="20.078" cy="10.517" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.078 12.766v7.6m15.944-2.868c0 1.577-1.304 2.868-2.899 2.868s-2.899-1.29-2.899-2.868v-1.864c0-1.577 1.305-2.868 2.9-2.868s2.898 1.29 2.898 2.868m0 4.732v-7.6"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M28.063 12.766v8.604c0 1.577-1.305 2.868-2.9 2.868c-.869 0-1.594-.287-2.029-.86"/><path d="M25.164 12.766c1.594 0 2.899 1.29 2.899 2.868v1.864c0 1.577-1.305 2.868-2.9 2.868s-2.898-1.29-2.898-2.868v-1.864c0-1.577 1.304-2.868 2.899-2.868Z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.475 36.795c-.434.86-1.45 1.434-2.464 1.434c-1.594 0-2.899-1.29-2.899-2.868v-1.864c0-1.577 1.305-2.867 2.9-2.867s2.898 1.29 2.898 2.867v1.004h-5.798M15.09 26.758v10.037c0 .86.58 1.434 1.45 1.434h.434m6.079.004V28.77c0-1.147.87-2.008 2.03-2.008c1.014 0 1.594.287 2.03.86m-5.653 3.012h4.059"/><circle cx="19.217" cy="28.38" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.217 30.629v7.6"/>`),
		g.Group(children),
	)
}