package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldCheckCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilShieldCheckCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M14.557 5.332a3.5 3.5 0 0 0-3.114 0L7.807 7.14C6.707 7.686 5.99 8.75 6 9.935c.016 1.776.207 4.132.949 5.688c.824 1.728 2.8 3.456 4.368 4.625c1.007.75 2.357.75 3.364 0c1.567-1.169 3.544-2.897 4.368-4.626c.742-1.555.933-3.91.949-5.687c.01-1.184-.705-2.25-1.806-2.796l-3.636-1.807Zm-3.56-.895a4.5 4.5 0 0 1 4.005 0l3.636 1.806c1.402.697 2.375 2.09 2.36 3.7c-.015 1.783-.2 4.337-1.045 6.11c-.94 1.972-3.103 3.826-4.673 4.997a3.794 3.794 0 0 1-4.56 0c-1.57-1.171-3.733-3.025-4.673-4.997C5.2 14.28 5.017 11.726 5 9.943c-.014-1.61.96-3.003 2.361-3.7l3.636-1.806Z"/><path d="M15.795 10.884c.212.178.24.494.065.707l-1.78 2.155a1.486 1.486 0 0 1-2.205.097l-1.208-1.22a.505.505 0 0 1-.001-.71a.495.495 0 0 1 .704.002l1.209 1.22c.207.21.548.194.735-.032l1.78-2.156a.495.495 0 0 1 .7-.063Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilShieldCheckCircleFilled0)"/></g>`),
		g.Group(children),
	)
}