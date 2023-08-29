package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Croissant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCroissant0"><g fill="none"><g stroke="#fff" stroke-width="4" clip-path="url(#ipSCroissant1)"><path stroke-linecap="round" stroke-linejoin="round" d="M7 29s-.044 3.748 1.456 8.248c.873 2.616 1.914 4.219 3.223 5.003c1.677 1.005 3.397-.98 4.379-2.67c.962-1.657 2.375-3.353 3.49-4.581M33 7.911c2.623.44 5.532 1.006 6.626 1.508c1.426.654 2.86 1.986 3.431 3.411c.531 1.324-.847 2.34-2.093 3.034c-1.405.783-2.869 2.01-4.065 3.136"/><path stroke-linecap="round" stroke-linejoin="round" d="M24.202 29.776L21.7 31.513a2.97 2.97 0 0 0-1.036 1.27v0a2.97 2.97 0 0 1-3.886 1.566L8.803 30.98a2.96 2.96 0 0 1-1.61-3.792l.45-1.169a11.112 11.112 0 0 1 4.576-5.486v0m18.501 3.975l1.39-1.504l.912-.987a3.312 3.312 0 0 1 1.125-.795v0a3.305 3.305 0 0 0 1.812-4.146l-2.852-8.01a2.797 2.797 0 0 0-3.259-1.789l-1.218.279A10.798 10.798 0 0 0 22.114 12v0"/><path d="M13.41 23.078a3.817 3.817 0 0 1-1.017-5.12a19.032 19.032 0 0 1 6.154-6.135a3.8 3.8 0 0 1 5.036.954l5.932 7.909a4 4 0 0 1-.913 5.682l-.24.167c-.301.21-.572.46-.804.743l-.684.832a4 4 0 0 1-5.416.716l-8.047-5.748Z"/></g><defs><clipPath id="ipSCroissant1"><path fill="#000" d="M0 0h48v48H0z"/></clipPath></defs></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCroissant0)"/>`),
		g.Group(children),
	)
}