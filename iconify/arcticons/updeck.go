package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Updeck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.168 32.318a21.504 21.504 0 1 1 40.98-12.211m.336 4.724A21.5 21.5 0 0 1 9.007 39.41"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.617 25.402c2.187 0 2.775 1.85 2.775 3.981c0 3.505-4.906 12.28-10.261 12.28c-1.458 0-2.495-.28-2.495-2.187s4.261-14.074 9.98-14.074Z"/><ellipse cx="40.734" cy="22.749" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.907" ry="3.651" transform="rotate(-7.549 40.734 22.749)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.255 19.13c-2.62 1.282-11.298 3.558-14.591 4.057a1.948 1.948 0 0 0-1.879 2.44c.252 1.009 1.065 1.513 2.523 1.233a44.408 44.408 0 0 1 14.906-.492"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.526 25.98q.091.838.128 1.707c.337 7.865-2.497 14.945-4.255 17.13M13.874 42.97a9.175 9.175 0 0 0 4.514-8.26c0-3.911-4.514-13.248-4.514-16.15s1.107-4.92 8.341-4.92a10.013 10.013 0 0 1 7.644 3.163"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.794 14.578c-4.074-1.863-9-3.313-14.967-.223c-5.846 3.028-3.53 12.399-3.53 12.399"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.194 10.717c.63-1.261 4.498-2.338 5.172 1.849m15.316 17.476c4.745.39 13.534-4.784 13.052-7.907c-.4-2.588-4.237-2.842-5.441-2.62M11.26 20.2l.127-1.316a.857.857 0 0 1 .932-.77l5.732.53m1.269-1.051h4.473a1.01 1.01 0 0 1 1.009 1.01v1.65a2.26 2.26 0 0 1-2.26 2.261H20.65a2.26 2.26 0 0 1-2.258-2.154l-.08-1.71a1.01 1.01 0 0 1 1.008-1.056Zm14.112 1.578l-.12-1.574a1.347 1.347 0 0 0-1.645-1.21l-3.93.904a1.347 1.347 0 0 0-1.01 1.617l.182.79a3.158 3.158 0 0 0 3.833 2.352l.74-.182a2.574 2.574 0 0 0 1.95-2.697Zm-8.631-.568h1.89"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.858 16.803c1.42.79 2.762 2.164 1.334 5.136"/><circle cx="21.771" cy="19.248" r="1.598" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.065" cy="18.603" r="1.335" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}