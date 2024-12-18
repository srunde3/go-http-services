package main

func (s *server) routes() {
	s.router.Get("/health", HandleHealthCheck)
	s.router.Get("/hi-mom", HandleHiMom)
	s.router.Get("/echo/{parameter}", HandleEchoUrlParameter)
}
