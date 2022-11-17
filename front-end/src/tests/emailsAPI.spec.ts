import { emailsSearch } from "../services/emailsAPI";
import { TermType } from "@/types/interface";

describe('emailsAPI test', () => {
    test('emailsSearch error', () => {
    let Term: TermType = {
        term: "",
      };
      expect(emailsSearch(Term,0,25)).toBe("You must introduce a term!")
    })
  })