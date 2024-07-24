package response

var (
	ErrorMessages []Error
)

func init() {
	ErrorMessages = []Error{
		{
			ID: RCUnknownError,
			Descriptions: map[string]interface{}{
				"EN": "General error has occurred",
				"ID": "Terjadi kesalahan umum",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCSystemError,
			Descriptions: map[string]interface{}{
				"EN": "System error has occurred",
				"ID": "Terjadi kesalahan pada sistem",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCDatabaseError,
			Descriptions: map[string]interface{}{
				"EN": "Database error has occurred",
				"ID": "Terjadi kesalahan pada database",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCFileSystemError,
			Descriptions: map[string]interface{}{
				"EN": "File system error has occurred",
				"ID": "Terjadi kesalahan pada sistem file",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCThirdPartySystemError,
			Descriptions: map[string]interface{}{
				"EN": "Third party system error has occurred",
				"ID": "Terjadi kesalahan pada sistem pihak ketiga",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCConnectionTimeout,
			Descriptions: map[string]interface{}{
				"EN": "Connection timeout",
				"ID": "Waktu koneksi habis atau terputus",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCDataNotFound,
			Descriptions: map[string]interface{}{
				"EN": "Data not found",
				"ID": "Data tidak ditemukan",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCDuplicateData,
			Descriptions: map[string]interface{}{
				"EN": "Data is already registered",
				"ID": "Data sudah pernah didaftarkan",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCImmutableData,
			Descriptions: map[string]interface{}{
				"EN": "Data cannot be edited",
				"ID": "Data tidak dapat diedit",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCNotAuthorizedAccess,
			Descriptions: map[string]interface{}{
				"EN": "Not authorized access",
				"ID": "Akses tidak diizinkan",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCIInvalidCredential,
			Descriptions: map[string]interface{}{
				"EN": "Invalid user credentials!",
				"ID": "Kredensial pengguna tidak valid!",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCUserIsLoggedIn,
			Descriptions: map[string]interface{}{
				"EN": "User is currently active",
				"ID": "User anda sedang aktif",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCInvalidLoginSession,
			Descriptions: map[string]interface{}{
				"EN": "Invalid login session",
				"ID": "Sesi login tidak valid",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCUnsupportedBodyType,
			Descriptions: map[string]interface{}{
				"EN": "Request body is not JSON format",
				"ID": "Format request body bukan JSON",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCMissingParameter,
			Descriptions: map[string]interface{}{
				"EN": "Missing required field(s)",
				"ID": "Input yang dibutuhkan tidak lengkap",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCInvalidInputFormat,
			Descriptions: map[string]interface{}{
				"EN": "Invalid input field format : ",
				"ID": "Format input salah : ",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCUploadFileFailed,
			Descriptions: map[string]interface{}{
				"EN": "Upload file failed",
				"ID": "Gagal mengupload file",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCOTPHasBeenSent,
			Descriptions: map[string]interface{}{
				"EN": "OTP has been sent, please check again",
				"ID": "OTP telah di kirim, silakan check kembali",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCOTPHasExpired,
			Descriptions: map[string]interface{}{
				"EN": "OTP expired",
				"ID": "OTP telah kedaluwarsa",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCOTPInvalid,
			Descriptions: map[string]interface{}{
				"EN": "OTP invalid",
				"ID": "OTP salah",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCPINInvalid,
			Descriptions: map[string]interface{}{
				"EN": "Invalid PIN",
				"ID": "PIN tidak valid",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCAccountNotFullySetup,
			Descriptions: map[string]interface{}{
				"EN": "Account is not fully setup",
				"ID": "Akun belum terkonfigurasi sepenuhnya",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCAccountDisabled,
			Descriptions: map[string]interface{}{
				"EN": "Account is disabled",
				"ID": "Akun dinonaktifkan",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCInvalidSignature,
			Descriptions: map[string]interface{}{
				"EN": "Invalid signature!",
				"ID": "Signature tidak valid!",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCOTPMaxAttempt,
			Descriptions: map[string]interface{}{
				"EN": "We're sorry you have entered the wrong OTP number as much as the limit that has been determined. Your account is now suspended until {{.timestamp}}.",
				"ID": "Mohon maaf Anda telah salah memasukkan nomor OTP sebanyak limit yang sudah ditentukan. Akun Anda sekarang ditangguhkan sampai dengan {{.timestamp}}.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCInvalidCode,
			Descriptions: map[string]interface{}{
				"EN": "Invalid Code",
				"ID": "Code yang anda masukan tidak valid",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCCodeHasExpired,
			Descriptions: map[string]interface{}{
				"EN": "Code is already expired",
				"ID": "Code yang anda masukan sudah tidak aktif",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCInvalidInputData,
			Descriptions: map[string]interface{}{
				"EN": "Input data is not valid",
				"ID": "Data yang anda masukan tidak valid",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCValidationProcessNotComplete,
			Descriptions: map[string]interface{}{
				"EN": "Validation process is not complete",
				"ID": "Proses validasi belum lengkap",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCUnregisteredDeviceID,
			Descriptions: map[string]interface{}{
				"EN": "Unregistered device ID",
				"ID": "Device ID belum terdaftar",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCCommunicateWithIbridge,
			Descriptions: map[string]interface{}{
				"EN": "Error communicate with IBridge",
				"ID": "Error komunikasi dengan IBridge",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCOTPAttemptBlockedPermanent,
			Descriptions: map[string]interface{}{
				"EN": "Your OTP request has been permanently blocked because it has reached the maximum trial use of the OTP code, please contact the call center for more information.",
				"ID": "Permintaan OTP anda telah di blokir permanen karena sudah mencapai maksimal percobaan penggunaan kode OTP, silahkan hubungi call center untuk informasi lebih lanjut.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCOTPAttemptBlockedByNumber,
			Descriptions: map[string]interface{}{
				"EN": "We're sorry you have entered the wrong OTP number as much as the limit that has been determined.",
				"ID": "Mohon maaf Anda telah salah memasukkan nomor OTP sebanyak limit yang sudah ditentukan.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCInternalServerError,
			Descriptions: map[string]interface{}{
				"EN": "An error occurred in the system, please try again!",
				"ID": "Terjadi kesalahan pada sistem, silahkan coba kembali!",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCAuthMaxAttempt,
			Descriptions: map[string]interface{}{
				"EN": "We're sorry you have entered the wrong authentication as much as the limit that has been determined. Your account is now suspended until {{.timestamp}}.",
				"ID": "Mohon maaf Anda telah salah melakukan percobaan autentikasi sebanyak limit yang sudah ditentukan. Akun Anda sekarang ditangguhkan sampai dengan {{.timestamp}}.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCAuthAttemptBlockedPermanent,
			Descriptions: map[string]interface{}{
				"EN": "Your authentication request has been permanently blocked because it has reached the maximum trial use of the authentication attempt, please contact the call center for more information.",
				"ID": "Permintaan autentikasi anda telah di blokir permanen karena sudah mencapai maksimal percobaan penggunaan autentikasi, silakan hubungi call center untuk informasi lebih lanjut.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCErrorGateway, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Error communication with gateway!",
				"ID": "Error komunikasi dengan gateway!",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCCardRegisteredDifference, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Card is already registered with different phone number!",
				"ID": "Kartu telah terdaftar dengan nomor telepon lain!",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCUserBlacklisted, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "You are blacklisted",
				"ID": "Kamu terdaftar dalam Daftar Hitam",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCInvalidPersonalData, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Invalid input personal data!",
				"ID": "Data personal yang anda masukan tidak valid!",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCNeedManualVerification, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Need manual verification admin",
				"ID": "Manual verifikasi admin",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCFaceUndetected, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Face undetected, please retry liveness test!",
				"ID": "Wajah tidak terdeteksi, harap mengulang livesness test!",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCTransactionMaxDailyLimit, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Your transaction has reach the daily limit",
				"ID": "Transaksi Anda telah mencapai limit harian",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCTransactionMaxLimit, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Your transaction amount exceed the maximum limit transaction",
				"ID": "Nominal transaksi Anda melebihi batas maksimum transaksi",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCInvalidExpiredCard, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Invalid expired card.",
				"ID": "Expired kartu tidak sesuai.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCInvalidATMCardNumber, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Invalid ATM card number.",
				"ID": "Nomor kartu ATM tidak sesuai.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCDataOnProcess, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Your data is on process",
				"ID": "Data anda sedang diproses",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCServiceNotFound, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Service not found",
				"ID": "Service tidak ditemukan",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCInvalidIPassport, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Invalid IPassport",
				"ID": "IPassport tidak valid",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCFaceDoesntMatch, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "The face doesn't match, please try again",
				"ID": "Wajah tidak cocok, silakan ulangi lagi.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCCantExecuteTask, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Can't execute task",
				"ID": "Tidak bisa mengeksekusi task",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCOTPAttemptBlocked,
			Descriptions: map[string]interface{}{
				"EN": "We're sorry your phone number is suspended until {{.timestamp}}",
				"ID": "Mohon maaf nomor telepon Anda ditangguhkan sampai {{.timestamp}}",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCUseOriginalKTP, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Please use the original e-KTP.",
				"ID": "Mohon menggunakan e-KTP asli.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCCantRegisterMB, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "You can not register new account through mobile banking. Please contact our call center for detailed information.",
				"ID": "Anda tidak dapat melakukan pembukaan rekening melalui kanal mobile banking. silakan menghubungi call center untuk informasi lebih lanjut.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCIDNumberAlreadyRegistered, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Your ID Number already registered in {{.bank_name}}.",
				"ID": "NIK anda sudah terdaftar mempunyai rekening pada {{.bank_name}}.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCImageCantBeRead, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Image can not be read by our system. Please retake a photo of your ID card.",
				"ID": "Gambar tidak terbaca oleh sistem kami. Mohon foto kembali KTP anda.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCIdentityNotMatchKTP, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"EN": "Identity data does not match the e-KTP.",
				"ID": "Data identitas tidak sesuai e-KTP.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCOTPMaxSend,
			Descriptions: map[string]interface{}{
				"EN": "Already limit for send otp. You can send otp on {{.timestamp}}.",
				"ID": "Sudah melebihi batas send otp. Kamu bisa melakukan lagi pada {{.timestamp}.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCOTPMaxResend,
			Descriptions: map[string]interface{}{
				"EN": "Already limit for send otp. You can resend otp on {{.timestamp}.",
				"ID": "Sudah melebihi batas resend otp. Kamu bisa melakukan lagi pada {{.timestamp}.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCOTPMaxVerify,
			Descriptions: map[string]interface{}{
				"EN": "OTP verification exceed trial limit. You can try verify otp on {{.timestamp}.",
				"ID": "Verifikasi otp melebihi batas percobaan. Kamu bisa melakukan lagi pada {{.timestamp}.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCBadRequest,
			Descriptions: map[string]interface{}{
				"EN": "Bad request.",
				"ID": "Data request tidak sesuai contract.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCConfirmationLinkSent, // id not found in folder response file error_message
			Descriptions: map[string]interface{}{
				"ID": "Link konfirmasi anda telah dikirimkan melalui email, silakan cek email anda",
				"EN": "Your confirmation link has been sent via email, please check your email",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCPINHasExpired,
			Descriptions: map[string]interface{}{
				"EN": "Key expired",
				"ID": "Kunci telah kedaluwarsa",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCRedisConnection,
			Descriptions: map[string]interface{}{
				"EN": "Error connection with Redis.",
				"ID": "Error komunikasi dengan Redis.",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCInsufficientBalance,
			Descriptions: map[string]interface{}{
				"EN": "Insufficient balance",
				"ID": "Saldo tidak mencukupi",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCTransactionNotAllowed,
			Descriptions: map[string]interface{}{
				"EN": "Transaction not allowed",
				"ID": "Transaksi tidak diizinkan",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCTransactionBelowMinimal,
			Descriptions: map[string]interface{}{
				"EN": "Your transaction amount is below the minimum amount",
				"ID": "Nominal transaksi anda dibawah minimal transaksi",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please contact call center",
				"ID": "Hubungi call center",
			},
		},
		{
			ID: RCTransactionAmountDifferent,
			Descriptions: map[string]interface{}{
				"EN": "Transaction amount is different from payment inquiry",
				"ID": "Nominal transaksi berbeda dengan saat cek pembayaran",
			},
			ProblemOwner:  "user",
			SeverityLevel: 5,
			WhatToDo: map[string]interface{}{
				"EN": "Please check your transaction amount",
				"ID": "Periksa kembali nominal transaksi yang anda masukan",
			},
		},
	}
}
