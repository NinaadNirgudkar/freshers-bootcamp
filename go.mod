module bootcamp/m

go 1.22.0

replace bootcamp/dayone => ./dayone

require bootcamp/daytwo v0.0.0-00010101000000-000000000000

require (
	bootcamp/dayone v0.0.0-00010101000000-000000000000 // indirect
	bootcamp/daythree v0.0.0-00010101000000-000000000000 // indirect
)

replace bootcamp/daytwo => ./daytwo

replace bootcamp/daythree => ./daythree
