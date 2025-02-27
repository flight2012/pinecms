import { BaseService, Service } from "/@/core";

@Service("document")
class SysDocument extends BaseService {
	select(params: any) {
		return this.request({
			url: "/select",
			method: "GET",
			params
		});
	}
}

export default SysDocument;
