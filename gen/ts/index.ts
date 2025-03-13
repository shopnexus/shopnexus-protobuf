import { AccountService } from "./pb/account/v1/service_pb"
import { LoginAdminRequestSchema } from "./pb/account/v1/account_pb"
import { create, toBinary, toJson } from "@bufbuild/protobuf"

async function main() {
	let user = create(LoginAdminRequestSchema, {
		username: "admin",
		password: "admin",
	})

	// const bytes = toBinary(LoginAdminRequestSchema, user)
	const json = toJson(LoginAdminRequestSchema, user)
	console.log(json)
}

main()
