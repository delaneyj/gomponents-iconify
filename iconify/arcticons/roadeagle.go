package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roadeagle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 25.467h39m-34.836 0v4.626m5.467-1.177h8.412l-.8 13.458H7.697l6.434-13.458zm25.205-3.449v4.626m-5.467-1.177h-8.412l.8 13.458h14.046l-6.434-13.458zm-8.893-19.46l.931-3.83l-3.07 2.103l-1.178-2.103l-1.177 2.103l-3.07-2.103l.931 3.83h6.633zm-10.929 2.89h1.514c1.514 0 2.734-1.514 3.49-1.514c2.076 0 7.347-.701 11.482 11.565H17.412s1.43-2.299 1.43-3.266s-2.131-1.991-3.267-1.991a2.36 2.36 0 0 0-1.794.953s-.35-1.864-.35-2.411s.616-3.336.616-3.336Z"/><circle cx="18.799" cy="13.902" r="1.178" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}