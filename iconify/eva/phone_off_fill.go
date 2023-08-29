package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPhoneOffFill0"><g id="evaPhoneOffFill1"><path id="evaPhoneOffFill2" fill="currentColor" d="M9.27 12.06a10.37 10.37 0 0 1-.8-1.42C9.71 10 9.72 10 9.85 9.85a1 1 0 0 0 .26-.92L8.74 3a1 1 0 0 0-.65-.72a3.79 3.79 0 0 0-.72-.18A3.94 3.94 0 0 0 6.6 2A4.6 4.6 0 0 0 2 6.6a15.33 15.33 0 0 0 3.27 9.46Zm12.67 4.58a4.34 4.34 0 0 0-.19-.73a1 1 0 0 0-.72-.65l-6-1.37a1 1 0 0 0-.92.26c-.14.13-.15.14-.8 1.38a10.88 10.88 0 0 1-1.41-.8l-4 4A15.33 15.33 0 0 0 17.4 22a4.6 4.6 0 0 0 4.6-4.6a4.77 4.77 0 0 0-.06-.76Zm-2.2-12.38a.89.89 0 0 0-1.26 0L4.26 18.48a.91.91 0 0 0-.26.63a.89.89 0 0 0 1.52.63L19.74 5.52a.89.89 0 0 0 0-1.26Z"/></g></g>`),
		g.Group(children),
	)
}