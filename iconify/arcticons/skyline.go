package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skyline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24.156" cy="12.405" r="6.905" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24.156" cy="35.595" r="6.905"/><circle cx="35.588" cy="23.959" r="6.905"/><circle cx="12.412" cy="23.959" r="6.905"/><path d="M21.094 6.271c-1.54 3.966-1.497 7.946-.069 12.256m5.415.179c-2.505-.938-4.139-3.582-4.005-6.479c.135-2.897 2.005-5.345 4.583-6m2.714 14.645c1.833.614 6.225 5.991 6.913 9.564M18.55 26.912c-1.833-.613-6.225-5.99-6.913-9.564m10.22 24.668c2.659-.79 4.471-3.482 4.374-6.499c-.096-3.017-2.076-5.564-4.778-6.148m5.904.25c1.54 3.966 1.497 7.738.069 12.049"/></g>`),
		g.Group(children),
	)
}