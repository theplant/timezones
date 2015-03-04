package timezones

type Zoneinfo struct {
	Offset         int
	TzName         string
	Name           string
	NameWithOffset string
}

var AllZones = []*Zoneinfo{

	{-10, "US/Hawaii", "Hawaii", "(GMT-10:00) Hawaii"},
	{-9, "US/Alaska", "Alaska", "(GMT-09:00) Alaska"},
	{-8, "US/Pacific", "Pacific Time (US & Canada)", "(GMT-08:00) Pacific Time (US & Canada)"},
	{-7, "US/Arizona", "Arizona", "(GMT-07:00) Arizona"},
	{-7, "US/Mountain", "Mountain Time (US & Canada)", "(GMT-07:00) Mountain Time (US & Canada)"},
	{-6, "US/Central", "Central Time (US & Canada)", "(GMT-06:00) Central Time (US & Canada)"},
	{-5, "US/Eastern", "Eastern Time (US & Canada)", "(GMT-05:00) Eastern Time (US & Canada)"},
	{-5, "US/East-Indiana", "Indiana (East)", "(GMT-05:00) Indiana (East)"},
	{-11, "US/Samoa", "International Date Line West", "(GMT-11:00) International Date Line West"},
	{-11, "Pacific/Midway", "Midway Island", "(GMT-11:00) Midway Island"},
	{-11, "US/Samoa", "Samoa", "(GMT-11:00) Samoa"},
	{-8, "America/Tijuana", "Tijuana", "(GMT-08:00) Tijuana"},
	{-7, "America/Chihuahua", "Chihuahua", "(GMT-07:00) Chihuahua"},
	{-7, "America/Mazatlan", "Mazatlan", "(GMT-07:00) Mazatlan"},
	{-6, "US/Central", "Central America", "(GMT-06:00) Central America"},
	{-6, "America/Mexico_City", "Guadalajara", "(GMT-06:00) Guadalajara"},
	{-6, "America/Mexico_City", "Mexico City", "(GMT-06:00) Mexico City"},
	{-6, "America/Monterrey", "Monterrey", "(GMT-06:00) Monterrey"},
	{-6, "Canada/Saskatchewan", "Saskatchewan", "(GMT-06:00) Saskatchewan"},
	{-5, "America/Bogota", "Bogota", "(GMT-05:00) Bogota"},
	{-5, "America/Lima", "Lima", "(GMT-05:00) Lima"},
	{-5, "America/Guayaquil", "Quito", "(GMT-05:00) Quito"},
	{-4, "America/Caracas", "Caracas", "(GMT-04:30) Caracas"},
	{-4, "Canada/Atlantic", "Atlantic Time (Canada)", "(GMT-04:00) Atlantic Time (Canada)"},
	{-4, "America/La_Paz", "La Paz", "(GMT-04:00) La Paz"},
	{-4, "America/Santiago", "Santiago", "(GMT-04:00) Santiago"},
	{-3, "Canada/Newfoundland", "Newfoundland", "(GMT-03:30) Newfoundland"},
	{-3, "America/Sao_Paulo", "Brasilia", "(GMT-03:00) Brasilia"},
	{-3, "America/Buenos_Aires", "Buenos Aires", "(GMT-03:00) Buenos Aires"},
	{-3, "America/Buenos_Aires", "Georgetown", "(GMT-03:00) Georgetown"},
	{-3, "America/Godthab", "Greenland", "(GMT-03:00) Greenland"},
	{-2, "America/Noronha", "Mid-Atlantic", "(GMT-02:00) Mid-Atlantic"},
	{-1, "Atlantic/Azores", "Azores", "(GMT-01:00) Azores"},
	{-1, "Atlantic/Cape_Verde", "Cape Verde Is.", "(GMT-01:00) Cape Verde Is."},
	{0, "Africa/Casablanca", "Casablanca", "(GMT+00:00) Casablanca"},
	{0, "Europe/Dublin", "Dublin", "(GMT+00:00) Dublin"},
	{0, "Europe/London", "Edinburgh", "(GMT+00:00) Edinburgh"},
	{0, "Europe/Lisbon", "Lisbon", "(GMT+00:00) Lisbon"},
	{0, "Europe/London", "London", "(GMT+00:00) London"},
	{0, "Africa/Monrovia", "Monrovia", "(GMT+00:00) Monrovia"},
	{0, "UTC", "UTC", "(GMT+00:00) UTC"},
	{1, "Europe/Amsterdam", "Amsterdam", "(GMT+01:00) Amsterdam"},
	{1, "Europe/Belgrade", "Belgrade", "(GMT+01:00) Belgrade"},
	{1, "Europe/Berlin", "Berlin", "(GMT+01:00) Berlin"},
	{1, "Europe/Berlin", "Bern", "(GMT+01:00) Bern"},
	{1, "Europe/Bratislava", "Bratislava", "(GMT+01:00) Bratislava"},
	{1, "Europe/Brussels", "Brussels", "(GMT+01:00) Brussels"},
	{1, "Europe/Budapest", "Budapest", "(GMT+01:00) Budapest"},
	{1, "Europe/Copenhagen", "Copenhagen", "(GMT+01:00) Copenhagen"},
	{1, "Europe/Ljubljana", "Ljubljana", "(GMT+01:00) Ljubljana"},
	{1, "Europe/Madrid", "Madrid", "(GMT+01:00) Madrid"},
	{1, "Europe/Paris", "Paris", "(GMT+01:00) Paris"},
	{1, "Europe/Prague", "Prague", "(GMT+01:00) Prague"},
	{1, "Europe/Rome", "Rome", "(GMT+01:00) Rome"},
	{1, "Europe/Sarajevo", "Sarajevo", "(GMT+01:00) Sarajevo"},
	{1, "Europe/Skopje	", "Skopje", "(GMT+01:00) Skopje"},
	{1, "Europe/Stockholm", "Stockholm", "(GMT+01:00) Stockholm"},
	{1, "Europe/Vienna", "Vienna", "(GMT+01:00) Vienna"},
	{1, "Europe/Warsaw", "Warsaw", "(GMT+01:00) Warsaw"},
	{1, "Africa/Lagos", "West Central Africa", "(GMT+01:00) West Central Africa"},
	{1, "Europe/Zagreb", "Zagreb", "(GMT+01:00) Zagreb"},
	{2, "Europe/Athens", "Athens", "(GMT+02:00) Athens"},
	{2, "Europe/Bucharest", "Bucharest", "(GMT+02:00) Bucharest"},
	{2, "Africa/Cairo", "Cairo", "(GMT+02:00) Cairo"},
	{2, "Africa/Harare", "Harare", "(GMT+02:00) Harare"},
	{2, "Europe/Helsinki", "Helsinki", "(GMT+02:00) Helsinki"},
	{2, "Europe/Istanbul", "Istanbul", "(GMT+02:00) Istanbul"},
	{2, "Asia/Jerusalem", "Jerusalem", "(GMT+02:00) Jerusalem"},
	{2, "Europe/Helsinki", "Kyiv", "(GMT+02:00) Kyiv"},
	{2, "Africa/Johannesburg", "Pretoria", "(GMT+02:00) Pretoria"},
	{2, "Europe/Riga", "Riga", "(GMT+02:00) Riga"},
	{2, "Europe/Sofia", "Sofia", "(GMT+02:00) Sofia"},
	{2, "Europe/Tallinn", "Tallinn", "(GMT+02:00) Tallinn"},
	{2, "Europe/Vilnius", "Vilnius", "(GMT+02:00) Vilnius"},
	{3, "Asia/Baghdad", "Baghdad", "(GMT+03:00) Baghdad"},
	{3, "Asia/Kuwait", "Kuwait", "(GMT+03:00) Kuwait"},
	{3, "Europe/Minsk", "Minsk", "(GMT+03:00) Minsk"},
	{3, "Europe/Moscow", "Moscow", "(GMT+03:00) Moscow"},
	{3, "Africa/Nairobi", "Nairobi", "(GMT+03:00) Nairobi"},
	{3, "Asia/Riyadh", "Riyadh", "(GMT+03:00) Riyadh"},
	{3, "Europe/Moscow", "St. Petersburg", "(GMT+03:00) St. Petersburg"},
	{3, "Asia/Tehran", "Tehran", "(GMT+03:30) Tehran"},
	{4, "Asia/Muscat", "Abu Dhabi", "(GMT+04:00) Abu Dhabi"},
	{4, "Asia/Baku", "Baku", "(GMT+04:00) Baku"},
	{4, "Asia/Muscat", "Muscat", "(GMT+04:00) Muscat"},
	{4, "Asia/Tbilisi", "Tbilisi", "(GMT+04:00) Tbilisi"},
	{4, "Europe/Volgograd", "Volgograd", "(GMT+04:00) Volgograd"},
	{4, "Asia/Yerevan", "Yerevan", "(GMT+04:00) Yerevan"},
	{4, "Asia/Kabul", "Kabul", "(GMT+04:30) Kabul"},
	{5, "Asia/Karachi", "Islamabad", "(GMT+05:00) Islamabad"},
	{5, "Asia/Karachi", "Karachi", "(GMT+05:00) Karachi"},
	{5, "Asia/Tashkent", "Tashkent", "(GMT+05:00) Tashkent"},
	{5, "Asia/Calcutta", "Chennai", "(GMT+05:30) Chennai"},
	{5, "Asia/Kolkata", "Kolkata", "(GMT+05:30) Kolkata"},
	{5, "Asia/Calcutta", "Mumbai", "(GMT+05:30) Mumbai"},
	{5, "Asia/Calcutta", "New Delhi", "(GMT+05:30) New Delhi"},
	{5, "Asia/Colombo", "Sri Jayawardenepura", "(GMT+05:30) Sri Jayawardenepura"},
	{5, "Asia/Kathmandu", "Kathmandu", "(GMT+05:45) Kathmandu"},
	{6, "Asia/Almaty", "Almaty", "(GMT+06:00) Almaty"},
	{6, "Asia/Dhaka", "Astana", "(GMT+06:00) Astana"},
	{6, "Asia/Dhaka", "Dhaka", "(GMT+06:00) Dhaka"},
	{6, "Asia/Yekaterinburg", "Ekaterinburg", "(GMT+06:00) Ekaterinburg"},
	{6, "Asia/Rangoon", "Rangoon", "(GMT+06:30) Rangoon"},
	{7, "Asia/Bangkok", "Bangkok", "(GMT+07:00) Bangkok"},
	{7, "Asia/Bangkok", "Hanoi", "(GMT+07:00) Hanoi"},
	{7, "Asia/Jakarta", "Jakarta", "(GMT+07:00) Jakarta"},
	{7, "Asia/Novosibirsk", "Novosibirsk", "(GMT+07:00) Novosibirsk"},
	{8, "Asia/Shanghai", "Beijing", "(GMT+08:00) Beijing"},
	{8, "Asia/Chongqing", "Chongqing", "(GMT+08:00) Chongqing"},
	{8, "Asia/Hong_Kong", "Hong Kong", "(GMT+08:00) Hong Kong"},
	{8, "Asia/Krasnoyarsk", "Krasnoyarsk", "(GMT+08:00) Krasnoyarsk"},
	{8, "Asia/Kuala_Lumpur", "Kuala Lumpur", "(GMT+08:00) Kuala Lumpur"},
	{8, "Australia/Perth", "Perth", "(GMT+08:00) Perth"},
	{8, "Asia/Singapore", "Singapore", "(GMT+08:00) Singapore"},
	{8, "Asia/Taipei", "Taipei", "(GMT+08:00) Taipei"},
	{8, "Asia/Ulaanbaatar", "Ulaan Bataar", "(GMT+08:00) Ulaan Bataar"},
	{8, "Asia/Urumqi", "Urumqi", "(GMT+08:00) Urumqi"},
	{9, "Asia/Irkutsk", "Irkutsk", "(GMT+09:00) Irkutsk"},
	{9, "Asia/Tokyo", "Osaka", "(GMT+09:00) Osaka"},
	{9, "Asia/Tokyo", "Sapporo", "(GMT+09:00) Sapporo"},
	{9, "Asia/Seou", "Seoul", "(GMT+09:00) Seoul"},
	{9, "Asia/Tokyo", "Tokyo", "(GMT+09:00) Tokyo"},
	{9, "Australia/Adelaide", "Adelaide", "(GMT+09:30) Adelaide"},
	{9, "Australia/Darwin", "Darwin", "(GMT+09:30) Darwin"},
	{10, "Australia/Brisbane", "Brisbane", "(GMT+10:00) Brisbane"},
	{10, "Australia/Canberra", "Canberra", "(GMT+10:00) Canberra"},
	{10, "Pacific/Guam", "Guam", "(GMT+10:00) Guam"},
	{10, "Australia/Hobart", "Hobart", "(GMT+10:00) Hobart"},
	{10, "Australia/Melbourne", "Melbourne", "(GMT+10:00) Melbourne"},
	{10, "Pacific/Port_Moresby", "Port Moresby", "(GMT+10:00) Port Moresby"},
	{10, "Australia/Sydney", "Sydney", "(GMT+10:00) Sydney"},
	{10, "Asia/Yakutsk", "Yakutsk", "(GMT+10:00) Yakutsk"},
	{11, "Asia/Magadan", "New Caledonia", "(GMT+11:00) New Caledonia"},
	{11, "Asia/Vladivostok", "Vladivostok", "(GMT+11:00) Vladivostok"},
	{12, "Pacific/Auckland", "Auckland", "(GMT+12:00) Auckland"},
	{12, "Pacific/Fiji", "Fiji", "(GMT+12:00) Fiji"},
	{12, "Asia/Kamchatka", "Kamchatka", "(GMT+12:00) Kamchatka"},
	{12, "Asia/Magadan", "Magadan", "(GMT+12:00) Magadan"},
	{12, "Pacific/Fiji", "Marshall Is.", "(GMT+12:00) Marshall Is."},
	{12, "Asia/Magadan", "Solomon Is.", "(GMT+12:00) Solomon Is."},
	{12, "Pacific/Auckland", "Wellington", "(GMT+12:00) Wellington"},
	{13, "Pacific/Tongatapu", "Nuku'alofa", "(GMT+13:00) Nuku'alofa"},
}
