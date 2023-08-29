package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenmrsLogoNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsOpenmrsLogoNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM15.768 15.768a11.622 11.622 0 0 1 8.244-3.414a11.62 11.62 0 0 1 8.226 3.398l.008-7.763A17.932 17.932 0 0 0 24.012 6c-2.968 0-5.784.808-8.253 2.08l.01 7.688Zm16.47 16.465a11.617 11.617 0 0 1-8.242 3.413c-3.212 0-6.12-1.298-8.228-3.398l-.009 7.763a17.937 17.937 0 0 0 8.236 1.99c2.968 0 5.769-.719 8.236-1.99l.006-7.778Zm-19.885-8.241c0 3.22 1.305 6.133 3.414 8.241l-7.778-.005A17.934 17.934 0 0 1 6 23.992c0-2.968.718-5.768 1.989-8.236l7.763.008a11.625 11.625 0 0 0-3.399 8.228Zm19.88-8.236a11.62 11.62 0 0 1 3.414 8.243c0 3.211-1.299 6.12-3.399 8.226l7.763.009a17.932 17.932 0 0 0 1.99-8.235a17.93 17.93 0 0 0-1.99-8.237l-7.778-.006Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsOpenmrsLogoNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}