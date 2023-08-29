package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="url(#f237id2)" d="M27.806 2.167a3.5 3.5 0 0 1 1.343 4.763L17.64 27.484a4 4 0 0 1-6.134 1.047l-8.723-7.684a3.5 3.5 0 1 1 4.627-5.253l5.473 4.821a.5.5 0 0 0 .766-.13L23.042 3.51a3.5 3.5 0 0 1 4.764-1.343Z"/><path fill="url(#f237id0)" d="M27.806 2.167a3.5 3.5 0 0 1 1.343 4.763L17.64 27.484a4 4 0 0 1-6.134 1.047l-8.723-7.684a3.5 3.5 0 1 1 4.627-5.253l5.473 4.821a.5.5 0 0 0 .766-.13L23.042 3.51a3.5 3.5 0 0 1 4.764-1.343Z"/><path fill="url(#f237id1)" d="M27.806 2.167a3.5 3.5 0 0 1 1.343 4.763L17.64 27.484a4 4 0 0 1-6.134 1.047l-8.723-7.684a3.5 3.5 0 1 1 4.627-5.253l5.473 4.821a.5.5 0 0 0 .766-.13L23.042 3.51a3.5 3.5 0 0 1 4.764-1.343Z"/><g filter="url(#f237id4)"><path stroke="url(#f237id3)" stroke-linecap="round" stroke-width="2" d="M25.47 6.865L15.207 25.01a1 1 0 0 1-1.526.264L6.47 19.028"/></g><defs><radialGradient id="f237id0" cx="0" cy="0" r="1" gradientTransform="matrix(-1.67772 2.6844 -2.70273 -1.68918 27.315 4.264)" gradientUnits="userSpaceOnUse"><stop stop-color="#79608A"/><stop offset="1" stop-color="#564065" stop-opacity="0"/></radialGradient><radialGradient id="f237id1" cx="0" cy="0" r="1" gradientTransform="matrix(-1.76161 -26.50833 41.10454 -2.7316 26.602 25.907)" gradientUnits="userSpaceOnUse"><stop stop-color="#984F70" stop-opacity="0"/><stop offset=".786" stop-color="#984F70" stop-opacity=".06"/><stop offset=".961" stop-color="#984F70"/></radialGradient><linearGradient id="f237id2" x1="28.783" x2="13.683" y1="7.116" y2="30.605" gradientUnits="userSpaceOnUse"><stop stop-color="#523E60"/><stop offset="1" stop-color="#0F080C"/></linearGradient><linearGradient id="f237id3" x1="25.092" x2="14.69" y1="8.878" y2="26.746" gradientUnits="userSpaceOnUse"><stop stop-color="#553F62"/><stop offset="1" stop-color="#29232C"/></linearGradient><filter id="f237id4" width="23.001" height="22.655" x="4.469" y="4.864" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feGaussianBlur result="effect1_foregroundBlur_18590_2020" stdDeviation=".5"/></filter></defs></g>`),
		g.Group(children),
	)
}