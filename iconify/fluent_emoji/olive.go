package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Olive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#6FA352" d="M25.267 6.724c5.8 5.8 6.35 14.64 1.23 19.76c-5.12 5.12-13.96 4.57-19.76-1.23c-5.8-5.8-6.35-14.64-1.23-19.75c5.12-5.11 13.96-4.57 19.76 1.22Z"/><path fill="url(#f1502id0)" d="M25.267 6.724c5.8 5.8 6.35 14.64 1.23 19.76c-5.12 5.12-13.96 4.57-19.76-1.23c-5.8-5.8-6.35-14.64-1.23-19.75c5.12-5.11 13.96-4.57 19.76 1.22Z"/><path fill="url(#f1502id1)" d="M25.267 6.724c5.8 5.8 6.35 14.64 1.23 19.76c-5.12 5.12-13.96 4.57-19.76-1.23c-5.8-5.8-6.35-14.64-1.23-19.75c5.12-5.11 13.96-4.57 19.76 1.22Z"/><path fill="url(#f1502id2)" d="M15.787 10.066c-2.119-.885-3.439-2.83-2.927-4.356c.504-1.517 2.632-2.033 4.759-1.158c2.127.876 3.439 2.83 2.926 4.357c-.512 1.527-2.63 2.042-4.758 1.157Z"/><path fill="#DB2956" d="M7.887 13.76c1.933 1.079 4.459.233 5.642-1.888c1.184-2.121.576-4.715-1.357-5.793C10.24 5 7.713 5.846 6.53 7.967c-1.184 2.121-.576 4.715 1.357 5.793Z"/><path fill="url(#f1502id3)" d="M7.887 13.76c1.933 1.079 4.459.233 5.642-1.888c1.184-2.121.576-4.715-1.357-5.793C10.24 5 7.713 5.846 6.53 7.967c-1.184 2.121-.576 4.715 1.357 5.793Z"/><defs><radialGradient id="f1502id0" cx="0" cy="0" r="1" gradientTransform="rotate(109.654 6.684 13.877) scale(20.4408 24.43)" gradientUnits="userSpaceOnUse"><stop stop-color="#91D060"/><stop offset=".458" stop-color="#6DAB32"/><stop offset="1" stop-color="#708D5D"/></radialGradient><radialGradient id="f1502id1" cx="0" cy="0" r="1" gradientTransform="matrix(2.87498 42.75004 -108.3363 7.28571 14 4.5)" gradientUnits="userSpaceOnUse"><stop stop-color="#596570" stop-opacity="0"/><stop offset=".311" stop-color="#5C6775" stop-opacity="0"/><stop offset=".698" stop-color="#616A7D"/></radialGradient><radialGradient id="f1502id2" cx="0" cy="0" r="1" gradientTransform="rotate(-104.412 14.463 .194) scale(11.551 12.5002)" gradientUnits="userSpaceOnUse"><stop stop-color="#C8E4BE" stop-opacity="0"/><stop offset=".437" stop-color="#C8E4BE" stop-opacity="0"/><stop offset=".832" stop-color="#C8E4BE"/></radialGradient><radialGradient id="f1502id3" cx="0" cy="0" r="1" gradientTransform="matrix(3.5625 -2.375 2.16411 3.24617 9.188 10.25)" gradientUnits="userSpaceOnUse"><stop stop-color="#E63B5F"/><stop offset="1" stop-color="#CC1138"/></radialGradient></defs></g>`),
		g.Group(children),
	)
}