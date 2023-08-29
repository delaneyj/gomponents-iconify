package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotorScooter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#ea5a47" d="M39.222 36.977s-5.963 17.859-15.684-4.109c-1.264-2.855-1.627-7.54-1.584-6.955l-5.9 12.136c-.376.782-1.239 1.484-.41 1.747c9.818 3.116 10.563 9.618 10.522 11.757c-.007.394 15.749.342 15.798 0c.453-3.167 3.464-17.566 19.856-6.224c1 .692-2.516-11.209-3.95-11.209c0 0-17.636.487-17.636 1.124l-1.012 1.733"/><circle cx="16.156" cy="50.893" r="5.476" fill="#d0cfce" transform="rotate(-1.335 16.175 50.952)"/><circle cx="51.398" cy="50.893" r="5.476" fill="#d0cfce"/><path fill="#a57939" d="M40.436 29.516h18.426v5.67H40.436z"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="16.156" cy="50.893" r="5.476" transform="rotate(-1.335 16.175 50.952)"/><circle cx="51.398" cy="50.893" r="5.476"/><path stroke-linecap="round" stroke-linejoin="round" d="M26.262 50.776a9.863 9.863 0 0 0-2.918-6.98c-3.89-3.874-10.2-3.874-14.09 0m52.443 6.98c0-5.48-4.442-9.921-9.922-9.921s-9.921 4.442-9.921 9.922m-1.418-21.261h18.426v5.67H40.436z"/><path stroke-linecap="round" stroke-linejoin="round" d="m14.923 40.855l8.505-17.009h5.669m-2.835 27.047h15.592m-1.418-15.708s-5.67 21.261-17.008 0m35.434 0s4.253 7.087 2.835 15.591"/></g>`),
		g.Group(children),
	)
}